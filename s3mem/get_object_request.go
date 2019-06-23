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

const opGetObject = "GetObject"

//GetObjectRequest ...
func (c *Client) GetObjectRequest(input *s3.GetObjectInput) s3.GetObjectRequest {
	if input == nil {
		input = &s3.GetObjectInput{}
	}
	output := &s3.GetObjectOutput{}
	operation := &aws.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}",
	}
	req := c.NewRequest(operation, input, output)
	return s3.GetObjectRequest{Request: req, Input: input}
}

func getObject(req *aws.Request) {
	S3MemService := GetS3MemService(req.Metadata.Endpoint)
	if !S3MemService.IsBucketExist(req.Params.(*s3.GetObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.GetObjectInput).Bucket, nil, nil)
	}
	obj, versionId, err := S3MemService.GetObject(req.Params.(*s3.GetObjectInput).Bucket, req.Params.(*s3.GetObjectInput).Key, req.Params.(*s3.GetObjectInput).VersionId)
	if err != nil {
		req.Error = err
		return
	}
	if obj == nil {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, req.Params.(*s3.GetObjectInput).Bucket, req.Params.(*s3.GetObjectInput).Key, req.Params.(*s3.GetObjectInput).VersionId)
		return
	}
	req.Data.(*s3.GetObjectOutput).Body = ioutil.NopCloser(bytes.NewReader(obj.Content))
	req.Data.(*s3.GetObjectOutput).VersionId = versionId
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.GetObjectOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
