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
	"time"

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const opCreateBucket = "CreateBucket"

//CreateBucketRequest ...
func (c *Client) CreateBucketRequest(input *s3.CreateBucketInput) s3.CreateBucketRequest {
	if input == nil {
		input = &s3.CreateBucketInput{}
	}
	output := &s3.CreateBucketOutput{
		Location: input.Bucket,
	}
	op := &aws.Operation{
		Name:       opCreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}
	req := c.NewRequest(op, input, output)
	return s3.CreateBucketRequest{Request: req, Input: input}
}

func createBucket(req *aws.Request) {
	S3MemService := GetS3MemService(req.Metadata.Endpoint)
	if S3MemService.IsBucketExist(req.Params.(*s3.CreateBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeBucketAlreadyExists, "", nil, req.Params.(*s3.CreateBucketInput).Bucket, nil, nil)
		return
	}
	tc := time.Now()
	bucket := &s3.Bucket{
		CreationDate: &tc,
		Name:         req.Params.(*s3.CreateBucketInput).Bucket,
	}
	S3MemService.CreateBucket(bucket)
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.CreateBucketOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
