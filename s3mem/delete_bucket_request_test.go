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
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

func TestDeleteBucketRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(S3MemEndpointsID, &s3.Bucket{Name: &bucketName})
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.DeleteBucketRequest(&s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)
	bucketGet := GetBucket(S3MemEndpointsID, &bucketName)
	assert.Nil(t, bucketGet)
}

func TestDeleteNotEmptyBucket(t *testing.T) {
	bucketName := strings.ToLower(t.Name())
	CreateBucket(S3MemEndpointsID, &s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	PutObject(S3MemEndpointsID, &bucketName, &objectKey, strings.NewReader(string("test content")))
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.DeleteBucketRequest(&s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.Error(t, err)
	assert.Implements(t, (*s3memerr.S3MemError)(nil), err)
	assert.Equal(t, s3memerr.ErrCodeBucketNotEmpty, err.(s3memerr.S3MemError).Code())
}
