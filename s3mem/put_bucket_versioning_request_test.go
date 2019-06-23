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

func TestPutBucketVersioningRequestMFAString(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})

	//Request a client
	client := New(S3MemTestConfig)
	mfa := "122334 13445"
	req := client.PutBucketVersioningRequest(&s3.PutBucketVersioningInput{
		Bucket: &bucketName,
		MFA:    &mfa,
	})
	//Send the request
	putBucketVersioningOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, putBucketVersioningOutput)
	mfaOut, _ := S3MemTestService.GetBucketVersioning(&bucketName)
	assert.NotNil(t, mfaOut)
	assert.Equal(t, mfa, *mfaOut)
}

func TestPutBucketVersioningRequestMFAEnable(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})

	//Request a client
	client := New(S3MemTestConfig)
	mfa := "122334 13445"
	req := client.PutBucketVersioningRequest(&s3.PutBucketVersioningInput{
		Bucket: &bucketName,
		MFA:    &mfa,
		VersioningConfiguration: &s3.VersioningConfiguration{
			MFADelete: s3.MFADeleteEnabled,
			Status:    s3.BucketVersioningStatusEnabled,
		},
	})
	//Send the request
	putBucketVersioningOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, putBucketVersioningOutput)
	_, versioningConfigurationOut := S3MemTestService.GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.MFADeleteEnabled, versioningConfigurationOut.MFADelete)
}

func TestPutBucketVersioningRequestMFADisabled(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})

	//Request a client
	client := New(S3MemTestConfig)
	mfa := "122334 13445"
	req := client.PutBucketVersioningRequest(&s3.PutBucketVersioningInput{
		Bucket: &bucketName,
		MFA:    &mfa,
		VersioningConfiguration: &s3.VersioningConfiguration{
			MFADelete: s3.MFADeleteDisabled,
		},
	})
	//Send the request
	putBucketVersioningOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, putBucketVersioningOutput)
	_, versioningConfigurationOut := S3MemTestService.GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.MFADeleteDisabled, versioningConfigurationOut.MFADelete)
}

func TestPutBucketVersioningRequestStatusEnabled(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})

	//Request a client
	client := New(S3MemTestConfig)
	mfa := "122334 13445"
	req := client.PutBucketVersioningRequest(&s3.PutBucketVersioningInput{
		Bucket: &bucketName,
		MFA:    &mfa,
		VersioningConfiguration: &s3.VersioningConfiguration{
			Status: s3.BucketVersioningStatusEnabled,
		},
	})
	//Send the request
	putBucketVersioningOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, putBucketVersioningOutput)
	_, versioningConfigurationOut := S3MemTestService.GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.BucketVersioningStatusEnabled, versioningConfigurationOut.Status)
}

func TestPutBucketVersioningRequestStatusSuspended(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})

	//Request a client
	client := New(S3MemTestConfig)
	mfa := "122334 13445"
	req := client.PutBucketVersioningRequest(&s3.PutBucketVersioningInput{
		Bucket: &bucketName,
		MFA:    &mfa,
		VersioningConfiguration: &s3.VersioningConfiguration{
			Status: s3.BucketVersioningStatusSuspended,
		},
	})
	//Send the request
	putBucketVersioningOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, putBucketVersioningOutput)
	_, versioningConfigurationOut := S3MemTestService.GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.BucketVersioningStatusSuspended, versioningConfigurationOut.Status)
}
