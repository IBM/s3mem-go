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
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const opListBuckets = "ListBuckets"

//ListBucketsRequest ...
func (c *Client) ListBucketsRequest(input *s3.ListBucketsInput) s3.ListBucketsRequest {
	if input == nil {
		input = &s3.ListBucketsInput{}
	}
	output := &s3.ListBucketsOutput{}
	op := &aws.Operation{
		Name:       opListBuckets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}
	req := c.NewRequest(op, input, output)
	return s3.ListBucketsRequest{Request: req, Input: input, Copy: c.ListBucketsRequest}
}

func listBuckets(req *aws.Request) {
	if req.Error != nil {
		return
	}
	req.Data.(*s3.ListBucketsOutput).Buckets = make([]s3.Bucket, 0)
	s3memBuckets := S3MemDatastores.Datastores[req.Metadata.Endpoint]
	var keys []string
	for k := range s3memBuckets.Buckets {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		req.Data.(*s3.ListBucketsOutput).Buckets = append(req.Data.(*s3.ListBucketsOutput).Buckets, *s3memBuckets.Buckets[k].Bucket)
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.ListBucketsOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
