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

const opDeleteBucket = "DeleteBucket"

//DeleteBucketRequest ...
func (c *Client) DeleteBucketRequest(input *s3.DeleteBucketInput) s3.DeleteBucketRequest {
	if input == nil {
		input = &s3.DeleteBucketInput{}
	}
	output := &s3.DeleteBucketOutput{}
	op := &aws.Operation{
		Name:       opDeleteBucket,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}",
	}
	req := c.NewRequest(op, input, output)
	return s3.DeleteBucketRequest{Request: req, Input: input}
}

func deleteBucket(req *aws.Request) {
	S3MemService := GetS3MemService(req.Metadata.Endpoint)
	if !S3MemService.IsBucketExist(req.Params.(*s3.DeleteBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteBucketInput).Bucket, nil, nil)
		return
	}
	if !S3MemService.IsBucketEmpty(req.Params.(*s3.DeleteBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3memerr.ErrCodeBucketNotEmpty, "", nil, req.Params.(*s3.DeleteBucketInput).Bucket, nil, nil)
		return
	}
	S3MemService.DeleteBucket(req.Params.(*s3.DeleteBucketInput).Bucket)
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.DeleteBucketOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
