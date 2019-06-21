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
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

const opDeleteObject = "DeleteObject"

//DeleteObjectRequest ...
func (c *Client) DeleteObjectRequest(input *s3.DeleteObjectInput) s3.DeleteObjectRequest {
	if input == nil {
		input = &s3.DeleteObjectInput{}
	}
	output := &s3.DeleteObjectOutput{}
	op := &aws.Operation{
		Name:       opDeleteObject,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}",
	}
	req := c.NewRequest(op, input, output)
	return s3.DeleteObjectRequest{Request: req, Input: input}
}

func deleteObject(req *aws.Request) {
	if !IsBucketExist(req.Metadata.Endpoint, req.Params.(*s3.DeleteObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteObjectInput).Bucket, nil, nil)
		return
	}
	if !IsObjectExist(req.Metadata.Endpoint, req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key, req.Params.(*s3.DeleteObjectInput).VersionId)
		return
	}
	deleteMarker, versionID, err := DeleteObject(req.Metadata.Endpoint, req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key, req.Params.(*s3.DeleteObjectInput).VersionId)
	if err != nil {
		req.Error = err
		return
	}
	req.Data.(*s3.DeleteObjectOutput).DeleteMarker = deleteMarker
	req.Data.(*s3.DeleteObjectOutput).VersionId = versionID
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.DeleteObjectOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
