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

const opPutBucketVersioning = "PutBucketVersioning"

//PutBucketVersioningRequest ...
func (c *Client) PutBucketVersioningRequest(input *s3.PutBucketVersioningInput) s3.PutBucketVersioningRequest {
	if input == nil {
		input = &s3.PutBucketVersioningInput{}
	}
	output := &s3.PutBucketVersioningOutput{}
	op := &aws.Operation{
		Name:       opPutBucketVersioning,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?versioning",
	}
	req := c.NewRequest(op, input, output)
	return s3.PutBucketVersioningRequest{Request: req, Input: input}
}

func putBucketVersioningBucketExists(req *aws.Request) {
}

func putBucketVersioning(req *aws.Request) {
	if !IsBucketExist(req.Metadata.Endpoint, req.Params.(*s3.PutBucketVersioningInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.PutBucketVersioningInput).Bucket, nil, nil)
		return
	}
	err := PutBucketVersioning(req.Metadata.Endpoint, req.Params.(*s3.PutBucketVersioningInput).Bucket, req.Params.(*s3.PutBucketVersioningInput).MFA, req.Params.(*s3.PutBucketVersioningInput).VersioningConfiguration)
	if err != nil {
		req.Error = s3memerr.NewError(s3memerr.ErrCodeIllegalVersioningConfigurationException, "", nil, req.Params.(*s3.PutBucketVersioningInput).Bucket, nil, nil)
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.PutBucketVersioningOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
