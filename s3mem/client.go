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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	s3.Client
}

func New(config aws.Config) *Client {
	svc := &Client{
		Client: *s3.New(config),
	}

	//set handlers
	svc.Handlers = Handlers()
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	//	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	// svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	// svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)
	svc.AddDebugHandlers()
	return svc
}

// AddDebugHandlers injects debug logging handlers into the service to log request
// debug information.
func (c *Client) AddDebugHandlers() {
	if !c.Config.LogLevel.AtLeast(aws.LogDebug) {
		return
	}

	c.Handlers.Send.PushFrontNamed(aws.NamedHandler{Name: "s3mem.client.LogRequest", Fn: logRequest})
	c.Handlers.Send.PushBackNamed(aws.NamedHandler{Name: "s3mem.client.LogResponse", Fn: logResponse})
}
