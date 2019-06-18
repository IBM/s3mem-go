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

func TestCreateBucketRequest(t *testing.T) {
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	bucketName := strings.ToLower(t.Name())
	req := client.CreateBucketRequest(&s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	//Send the request
	createBucketsOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	bucketGet := GetBucket(&bucketName)
	assert.NotNil(t, bucketGet)
	assert.Equal(t, bucketName, *bucketGet.Name)
	assert.Equal(t, bucketName, *createBucketsOutput.Location)
	assert.NotNil(t, bucketGet.CreationDate)
}

func TestCreateBucketRequestBucketAlreadyExists(t *testing.T) {
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	bucketName := strings.ToLower(t.Name())
	req := client.CreateBucketRequest(&s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	ctx := context.Background()
	//Send the request
	_, err := req.Send(ctx)
	//Assert the result
	assert.NoError(t, err)
	//Send the request
	_, err = req.Send(ctx)
	//Assert the result
	assert.Error(t, err)
	assert.Implements(t, (*s3memerr.S3MemError)(nil), err)
	assert.Equal(t, s3memerr.ErrCodeBucketAlreadyExists, err.(s3memerr.S3MemError).Code())
}
