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

func TestGetBucketVersioningRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})

	mfa := "122334 13445"

	//Add a VersionConfig
	S3MemTestService.PutBucketVersioning(&bucketName, &mfa, &s3.VersioningConfiguration{
		MFADelete: s3.MFADeleteEnabled,
		Status:    s3.BucketVersioningStatusEnabled,
	})
	//Request a client
	client := New(S3MemTestConfig)

	//Create request
	req := client.GetBucketVersioningRequest(&s3.GetBucketVersioningInput{
		Bucket: &bucketName,
	})

	//Send the request
	getBucketVersioningOut, err := req.Send(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, s3.MFADeleteStatusEnabled, getBucketVersioningOut.MFADelete)
	assert.Equal(t, s3.BucketVersioningStatusEnabled, getBucketVersioningOut.Status)
}
