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

	"github.com/IBM/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const opDeleteObjects = "DeleteObjects"

//DeleteObjectsRequest ...
func (c *Client) DeleteObjectsRequest(input *s3.DeleteObjectsInput) s3.DeleteObjectsRequest {
	if input == nil {
		input = &s3.DeleteObjectsInput{}
	}
	output := &s3.DeleteObjectsOutput{
		Deleted: make([]s3.DeletedObject, 0),
		Errors:  make([]s3.Error, 0),
	}
	op := &aws.Operation{
		Name:       opDeleteObjects,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}?delete",
	}
	req := c.NewRequest(op, input, output)
	return s3.DeleteObjectsRequest{Request: req, Input: input}
}

func deleteObjects(req *aws.Request) {
	S3MemService := GetS3MemService(req.Metadata.Endpoint)
	if !S3MemService.IsBucketExist(req.Params.(*s3.DeleteObjectsInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteObjectsInput).Bucket, nil, nil)
		return
	}
	for _, obj := range req.Params.(*s3.DeleteObjectsInput).Delete.Objects {
		deleteMarker, versionID, err := S3MemService.DeleteObject(req.Params.(*s3.DeleteObjectsInput).Bucket, obj.Key, obj.VersionId)
		if err != nil {
			req.Data.(*s3.DeleteObjectsOutput).Errors = append(req.Data.(*s3.DeleteObjectsOutput).Errors, err.Convert2S3Error(obj.Key, obj.VersionId))
		}
		req.Data.(*s3.DeleteObjectsOutput).Deleted = append(req.Data.(*s3.DeleteObjectsOutput).Deleted, s3.DeletedObject{
			DeleteMarker:          deleteMarker,
			DeleteMarkerVersionId: versionID,
			VersionId:             obj.VersionId,
			Key:                   obj.Key,
		})
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.DeleteObjectsOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
