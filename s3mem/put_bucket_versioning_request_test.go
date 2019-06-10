/*
###############################################################################
# Licensed Materials - Property of IBM Copyright IBM Corporation 2017, 2019. All Rights Reserved.
# U.S. Government Users Restricted Rights - Use, duplication or disclosure restricted by GSA ADP
# Schedule Contract with IBM Corp.
#
# Contributors:
#  IBM Corporation - initial API and implementation
###############################################################################
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
	CreateBucket(&s3.Bucket{Name: &bucketName})

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
	mfaOut, _ := GetBucketVersioning(&bucketName)
	assert.NotNil(t, mfaOut)
	assert.Equal(t, mfa, *mfaOut)
}

func TestPutBucketVersioningRequestMFAEnable(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})

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
	_, versioningConfigurationOut := GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.MFADeleteEnabled, versioningConfigurationOut.MFADelete)
}

func TestPutBucketVersioningRequestMFADisabled(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})

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
	_, versioningConfigurationOut := GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.MFADeleteDisabled, versioningConfigurationOut.MFADelete)
}

func TestPutBucketVersioningRequestStatusEnabled(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})

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
	_, versioningConfigurationOut := GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.BucketVersioningStatusEnabled, versioningConfigurationOut.Status)
}

func TestPutBucketVersioningRequestStatusSuspended(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})

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
	_, versioningConfigurationOut := GetBucketVersioning(&bucketName)
	assert.NotNil(t, versioningConfigurationOut)
	assert.Equal(t, s3.BucketVersioningStatusSuspended, versioningConfigurationOut.Status)
}
