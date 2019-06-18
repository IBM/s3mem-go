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

const opGetBucketVersioning = "GetBucketVersioning"

//GetBucketVersioningRequest ...
func (c *Client) GetBucketVersioningRequest(input *s3.GetBucketVersioningInput) s3.GetBucketVersioningRequest {
	if input == nil {
		input = &s3.GetBucketVersioningInput{}
	}
	output := &s3.GetBucketVersioningOutput{}
	op := &aws.Operation{
		Name:       opGetBucketVersioning,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}?versioning",
	}
	req := c.NewRequest(op, input, output)
	return s3.GetBucketVersioningRequest{Request: req, Input: input}
}

func getBucketVersioning(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.GetBucketVersioningInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.GetBucketVersioningInput).Bucket, nil, nil)
		return
	}
	_, obj := GetBucketVersioning(req.Params.(*s3.GetBucketVersioningInput).Bucket)
	if obj == nil {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.GetBucketVersioningInput).Bucket, nil, nil)
		return
	}
	switch obj.MFADelete {
	case s3.MFADeleteEnabled:
		req.Data.(*s3.GetBucketVersioningOutput).MFADelete = s3.MFADeleteStatusEnabled
	case s3.MFADeleteDisabled:
		req.Data.(*s3.GetBucketVersioningOutput).MFADelete = s3.MFADeleteStatusDisabled
	default:
	}
	req.Data.(*s3.GetBucketVersioningOutput).Status = obj.Status
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.GetBucketVersioningOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
