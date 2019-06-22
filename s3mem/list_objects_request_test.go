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
)

func TestListObjectssRequest(t *testing.T) {
	//Need to lock for testing as tests are running concurrently
	//and meanwhile another running test could change the stored buckets
	S3Store.S3MemServices[S3MemEndpointsID].Mux.Lock()
	defer S3Store.S3MemServices[S3MemEndpointsID].Mux.Unlock()

	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey1 := "1-my-object"
	content1 := "test content 1"
	S3MemTestService.PutObject(&bucketName, &objectKey1, strings.NewReader(string(content1)))
	objectKey2 := "2-my-object"
	content2 := "test content 2"
	S3MemTestService.PutObject(&bucketName, &objectKey2, strings.NewReader(string(content2)))

	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.ListObjectsRequest(&s3.ListObjectsInput{
		Bucket: &bucketName,
	})
	//Send the request
	listObjectsOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 2, len(listObjectsOutput.Contents))
	//Create the request
	prefix := "1"
	req = client.ListObjectsRequest(&s3.ListObjectsInput{
		Bucket: &bucketName,
		Prefix: &prefix,
	})
	//Send the request
	listObjectsOutput, err = req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 1, len(listObjectsOutput.Contents))
}
