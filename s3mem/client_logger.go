/*
################################################################################
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
*/
package s3mem

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http/httputil"

	"github.com/aws/aws-sdk-go-v2/aws"
)

const logReqMsg = `DEBUG: Request %s/%s Details:
---[ REQUEST POST-SIGN ]-----------------------------
%s
-----------------------------------------------------`

const logReqErrMsg = `DEBUG ERROR: Request %s/%s:
---[ REQUEST DUMP ERROR ]-----------------------------
%s
------------------------------------------------------`

type logWriter struct {
	// Logger is what we will use to log the payload of a response.
	Logger aws.Logger
	// buf stores the contents of what has been read
	buf *bytes.Buffer
}

func (logger *logWriter) Write(b []byte) (int, error) {
	return logger.buf.Write(b)
}

type teeReaderCloser struct {
	// io.Reader will be a tee reader that is used during logging.
	// This structure will read from a body and write the contents to a logger.
	io.Reader
	// Source is used just to close when we are done reading.
	Source io.ReadCloser
}

func (reader *teeReaderCloser) Close() error {
	return reader.Source.Close()
}

func logRequest(r *aws.Request) {
	logBody := r.Config.LogLevel.Matches(aws.LogDebugWithHTTPBody)
	dumpedBody, err := httputil.DumpRequestOut(r.HTTPRequest, false)
	if err != nil {
		r.Config.Logger.Log(fmt.Sprintf(logReqErrMsg, r.Metadata.ServiceName, r.Operation.Name, err))
		return
	}

	b, err := ioutil.ReadAll(r.HTTPRequest.Body)
	if err != nil {
		r.Config.Logger.Log(fmt.Sprintf(logRespErrMsg, r.Metadata.ServiceName, r.Operation.Name, err))
		return
	}
	r.Config.Logger.Log(fmt.Sprintf(logReqMsg, r.Metadata.ServiceName, r.Operation.Name, string(dumpedBody)))
	if r.Config.LogLevel.Matches(aws.LogDebugWithHTTPBody) {
		r.Config.Logger.Log(fmt.Sprintf("\n%s", string(b)))
	}

	if logBody {
		// Reset the request body because dumpRequest will re-wrap the r.HTTPRequest's
		// Body as a NoOpCloser and will not be reset after read by the HTTP
		// client reader.
		r.ResetBody()
	}

}

const logRespMsg = `DEBUG: Response %s/%s Details:
---[ RESPONSE ]--------------------------------------
%s
-----------------------------------------------------`

const logRespErrMsg = `DEBUG ERROR: Response %s/%s:
---[ RESPONSE DUMP ERROR ]-----------------------------
%s
-----------------------------------------------------`

func logResponse(r *aws.Request) {
	if r.Data == nil {
		r.Config.Logger.Log(fmt.Sprintf(logRespErrMsg,
			r.Metadata.ServiceName, r.Operation.Name, "request's Data is nil"))
		return
	}

	handlerFn := func(req *aws.Request) {
		body, err := httputil.DumpRequestOut(r.HTTPRequest, false)
		if err != nil {
			r.Config.Logger.Log(fmt.Sprintf(logRespErrMsg, req.Metadata.ServiceName, req.Operation.Name, err))
			return
		}
		b, err := ioutil.ReadAll(req.HTTPResponse.Body)
		if err != nil {
			r.Config.Logger.Log(fmt.Sprintf(logRespErrMsg, req.Metadata.ServiceName, req.Operation.Name, err))
			return
		}
		r.Config.Logger.Log(fmt.Sprintf(logRespMsg, req.Metadata.ServiceName, req.Operation.Name, string(body)))
		if req.Config.LogLevel.Matches(aws.LogDebugWithHTTPBody) {
			r.Config.Logger.Log(fmt.Sprintf("\n%s", string(b)))
		}
	}

	const handlerName = "s3mem.client.LogResponse.ResponseBody"

	r.Handlers.Unmarshal.SetBackNamed(aws.NamedHandler{
		Name: handlerName, Fn: handlerFn,
	})
	r.Handlers.UnmarshalError.SetBackNamed(aws.NamedHandler{
		Name: handlerName, Fn: handlerFn,
	})
}
