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
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/IBM/s3mem-go/s3mem/s3memerr"
)

const opListObjects = "ListObjects"

//ListObjectsRequest ...
func (c *Client) ListObjectsRequest(input *s3.ListObjectsInput) s3.ListObjectsRequest {
	if input == nil {
		input = &s3.ListObjectsInput{}
	}
	output := &s3.ListObjectsOutput{}
	op := &aws.Operation{
		Name:       opListObjects,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}",
		Paginator: &aws.Paginator{
			InputTokens:     []string{"Marker"},
			OutputTokens:    []string{"NextMarker || Contents[-1].Key"},
			LimitToken:      "MaxKeys",
			TruncationToken: "IsTruncated",
		},
	}
	req := c.NewRequest(op, input, output)
	return s3.ListObjectsRequest{Request: req, Input: input, Copy: c.ListObjectsRequest}
}

func listObjects(req *aws.Request) {
	S3MemService := GetS3MemService(req.Metadata.Endpoint)
	if !S3MemService.IsBucketExist(req.Params.(*s3.ListObjectsInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.ListObjectsInput).Bucket, nil, nil)
		return
	}
	req.Data.(*s3.ListObjectsOutput).Contents = make([]s3.Object, 0)
	bucket := req.Params.(*s3.ListObjectsInput).Bucket
	prefix := req.Params.(*s3.ListObjectsInput).Prefix
	s3memBuckets := S3Store.S3MemServices[req.Metadata.Endpoint]
	var keys []string
	for k := range s3memBuckets.Buckets[*bucket].Objects {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		obj := s3memBuckets.Buckets[*bucket].Objects[k]
		if prefix != nil {
			if strings.HasPrefix(*obj.VersionedObjects[0].Object.Key, *prefix) {
				req.Data.(*s3.ListObjectsOutput).Contents = append(req.Data.(*s3.ListObjectsOutput).Contents, *obj.VersionedObjects[0].Object)
			}
		} else {
			req.Data.(*s3.ListObjectsOutput).Contents = append(req.Data.(*s3.ListObjectsOutput).Contents, *obj.VersionedObjects[0].Object)
		}
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.ListObjectsOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
