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

func TestListBucketsRequest(t *testing.T) {
	//Need to lock for testing as tests are running concurrently
	//and meanwhile another running test could change the stored buckets
	S3MemTestService.Lock()
	defer S3MemTestService.Unlock()
	l := len(S3Store.S3MemServices[S3MemEndpointsID].Buckets)
	//Adding bucket directly in mem to prepare the test.
	bucket0 := strings.ToLower(t.Name() + "0")
	bucket1 := strings.ToLower(t.Name() + "1")
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucket0})
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucket1})
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.ListBucketsRequest(&s3.ListBucketsInput{})
	//Send the request
	listBucketsOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, l+2, len(listBucketsOutput.Buckets))
}
