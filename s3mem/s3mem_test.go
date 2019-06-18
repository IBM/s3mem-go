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
	"testing"

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws/endpoints"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

var S3MemTestConfig aws.Config

func init() {
	S3MemTestConfig = aws.Config{
		Region:      endpoints.EuWest1RegionID,
		LogLevel:    aws.LogDebugWithHTTPBody,
		Logger:      aws.NewDefaultLogger(),
		Credentials: aws.NewStaticCredentialsProvider("fake", "fake", ""),
	}
}

func TestNewClient(t *testing.T) {
	S3MemBuckets.Mux.Lock()
	defer S3MemBuckets.Mux.Unlock()
	l := len(S3MemBuckets.Buckets)
	client := New(S3MemTestConfig)
	//Create the request
	req := client.ListBucketsRequest(&s3.ListBucketsInput{})
	//Send the request
	listBucketsOutput, err := req.Send(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, l, len(listBucketsOutput.Buckets))
}

func TestNotImplemented(t *testing.T) {
	//Request a client
	client := New(S3MemTestConfig)
	input := &s3.AbortMultipartUploadInput{}
	req := client.AbortMultipartUploadRequest(input)
	assert.Error(t, req.Error)
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//AbortMultipartUploadRequest ...
func TestAbortMultipartUploadRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.AbortMultipartUploadRequest(&s3.AbortMultipartUploadInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//CompleteMultipartUploadRequest ...
func TestCompleteMultipartUploadRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.CompleteMultipartUploadRequest(&s3.CompleteMultipartUploadInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//CreateMultipartUploadRequest ...
func TestCreateMultipartUploadRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.CreateMultipartUploadRequest(&s3.CreateMultipartUploadInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketAnalyticsConfigurationRequest ...
func TestDeleteBucketAnalyticsConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketAnalyticsConfigurationRequest(&s3.DeleteBucketAnalyticsConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketCorsRequest ...
func TestDeleteBucketCorsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketCorsRequest(&s3.DeleteBucketCorsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketEncryptionRequest ...
func TestDeleteBucketEncryptionRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketEncryptionRequest(&s3.DeleteBucketEncryptionInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketInventoryConfigurationRequest ...
func TestDeleteBucketInventoryConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketInventoryConfigurationRequest(&s3.DeleteBucketInventoryConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketLifecycleRequest ...
func TestDeleteBucketLifecycleRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketLifecycleRequest(&s3.DeleteBucketLifecycleInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketMetricsConfigurationRequest ...
func TestDeleteBucketMetricsConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketMetricsConfigurationRequest(&s3.DeleteBucketMetricsConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketPolicyRequest ...
func TestDeleteBucketPolicyRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketPolicyRequest(&s3.DeleteBucketPolicyInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketReplicationRequest ...
func TestDeleteBucketReplicationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketReplicationRequest(&s3.DeleteBucketReplicationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketTaggingRequest ...
func TestDeleteBucketTaggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketTaggingRequest(&s3.DeleteBucketTaggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteBucketWebsiteRequest ...
func TestDeleteBucketWebsiteRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteBucketWebsiteRequest(&s3.DeleteBucketWebsiteInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeleteObjectTaggingRequest ...
func TestDeleteObjectTaggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeleteObjectTaggingRequest(&s3.DeleteObjectTaggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//DeletePublicAccessBlockRequest ...
func TestDeletePublicAccessBlockRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.DeletePublicAccessBlockRequest(&s3.DeletePublicAccessBlockInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketAccelerateConfigurationRequest ...
func TestGetBucketAccelerateConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketAccelerateConfigurationRequest(&s3.GetBucketAccelerateConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketAclRequest ...
func TestGetBucketAclRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketAclRequest(&s3.GetBucketAclInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketAnalyticsConfigurationRequest ...
func TestGetBucketAnalyticsConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketAnalyticsConfigurationRequest(&s3.GetBucketAnalyticsConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketCorsRequest ...
func TestGetBucketCorsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketCorsRequest(&s3.GetBucketCorsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketEncryptionRequest ...
func TestGetBucketEncryptionRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketEncryptionRequest(&s3.GetBucketEncryptionInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketInventoryConfigurationRequest ...
func TestGetBucketInventoryConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketInventoryConfigurationRequest(&s3.GetBucketInventoryConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketLifecycleRequest ...
func TestGetBucketLifecycleRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketLifecycleRequest(&s3.GetBucketLifecycleInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketLifecycleConfigurationRequest ...
func TestGetBucketLifecycleConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketLifecycleConfigurationRequest(&s3.GetBucketLifecycleConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketLocationRequest ...
func TestGetBucketLocationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketLocationRequest(&s3.GetBucketLocationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketLoggingRequest ...
func TestGetBucketLoggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketLoggingRequest(&s3.GetBucketLoggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketMetricsConfigurationRequest ...
func TestGetBucketMetricsConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketMetricsConfigurationRequest(&s3.GetBucketMetricsConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketNotificationRequest ...
func TestGetBucketNotificationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketNotificationRequest(&s3.GetBucketNotificationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketNotificationConfigurationRequest ...
func TestGetBucketNotificationConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketNotificationConfigurationRequest(&s3.GetBucketNotificationConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketPolicyRequest ...
func TestGetBucketPolicyRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketPolicyRequest(&s3.GetBucketPolicyInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketPolicyStatusRequest ...
func TestGetBucketPolicyStatusRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketPolicyStatusRequest(&s3.GetBucketPolicyStatusInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketReplicationRequest ...
func TestGetBucketReplicationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketReplicationRequest(&s3.GetBucketReplicationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketRequestPaymentRequest ...
func TestGetBucketRequestPaymentRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketRequestPaymentRequest(&s3.GetBucketRequestPaymentInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketTaggingRequest ...
func TestGetBucketTaggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketTaggingRequest(&s3.GetBucketTaggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetBucketWebsiteRequest ...
func TestGetBucketWebsiteRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetBucketWebsiteRequest(&s3.GetBucketWebsiteInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetObjectAclRequest ...
func TestGetObjectAclRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetObjectAclRequest(&s3.GetObjectAclInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetObjectLegalHoldRequest ...
func TestGetObjectLegalHoldRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetObjectLegalHoldRequest(&s3.GetObjectLegalHoldInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetObjectLockConfigurationRequest ...
func TestGetObjectLockConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetObjectLockConfigurationRequest(&s3.GetObjectLockConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetObjectRetentionRequest ...
func TestGetObjectRetentionRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetObjectRetentionRequest(&s3.GetObjectRetentionInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetObjectTaggingRequest ...
func TestGetObjectTaggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetObjectTaggingRequest(&s3.GetObjectTaggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetObjectTorrentRequest ...
func TestGetObjectTorrentRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetObjectTorrentRequest(&s3.GetObjectTorrentInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//GetPublicAccessBlockRequest ...
func TestGetPublicAccessBlockRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.GetPublicAccessBlockRequest(&s3.GetPublicAccessBlockInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//HeadBucketRequest ...
func TestHeadBucketRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.HeadBucketRequest(&s3.HeadBucketInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//HeadObjectRequest ...
func TestHeadObjectRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.HeadObjectRequest(&s3.HeadObjectInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//ListBucketAnalyticsConfigurationsRequest ...
func TestListBucketAnalyticsConfigurationsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.ListBucketAnalyticsConfigurationsRequest(&s3.ListBucketAnalyticsConfigurationsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//ListBucketInventoryConfigurationsRequest ...
func TestListBucketInventoryConfigurationsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.ListBucketInventoryConfigurationsRequest(&s3.ListBucketInventoryConfigurationsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//ListBucketMetricsConfigurationsRequest ...
func TestListBucketMetricsConfigurationsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.ListBucketMetricsConfigurationsRequest(&s3.ListBucketMetricsConfigurationsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//ListMultipartUploadsRequest ...
func TestListMultipartUploadsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.ListMultipartUploadsRequest(&s3.ListMultipartUploadsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//ListObjectVersionsRequest ...
func TestListObjectVersionsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.ListObjectVersionsRequest(&s3.ListObjectVersionsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//ListObjectsV2Request ...
func TestListObjectsV2Request(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.ListObjectsV2Request(&s3.ListObjectsV2Input{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//ListPartsRequest ...
func TestListPartsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.ListPartsRequest(&s3.ListPartsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketAccelerateConfigurationRequest ...
func TestPutBucketAccelerateConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketAccelerateConfigurationRequest(&s3.PutBucketAccelerateConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketAclRequest ...
func TestPutBucketAclRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketAclRequest(&s3.PutBucketAclInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketAnalyticsConfigurationRequest ...
func TestPutBucketAnalyticsConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketAnalyticsConfigurationRequest(&s3.PutBucketAnalyticsConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketCorsRequest ...
func TestPutBucketCorsRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketCorsRequest(&s3.PutBucketCorsInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketEncryptionRequest ...
func TestPutBucketEncryptionRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketEncryptionRequest(&s3.PutBucketEncryptionInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketInventoryConfigurationRequest ...
func TestPutBucketInventoryConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketInventoryConfigurationRequest(&s3.PutBucketInventoryConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketLifecycleRequest ...
func TestPutBucketLifecycleRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketLifecycleRequest(&s3.PutBucketLifecycleInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketLifecycleConfigurationRequest ...
func TestPutBucketLifecycleConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketLifecycleConfigurationRequest(&s3.PutBucketLifecycleConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketLoggingRequest ...
func TestPutBucketLoggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketLoggingRequest(&s3.PutBucketLoggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketMetricsConfigurationRequest ...
func TestPutBucketMetricsConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketMetricsConfigurationRequest(&s3.PutBucketMetricsConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketNotificationRequest ...
func TestPutBucketNotificationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketNotificationRequest(&s3.PutBucketNotificationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketNotificationConfigurationRequest ...
func TestPutBucketNotificationConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketNotificationConfigurationRequest(&s3.PutBucketNotificationConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketPolicyRequest ...
func TestPutBucketPolicyRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketPolicyRequest(&s3.PutBucketPolicyInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketReplicationRequest ...
func TestPutBucketReplicationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketReplicationRequest(&s3.PutBucketReplicationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketRequestPaymentRequest ...
func TestPutBucketRequestPaymentRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketRequestPaymentRequest(&s3.PutBucketRequestPaymentInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketTaggingRequest ...
func TestPutBucketTaggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketTaggingRequest(&s3.PutBucketTaggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutBucketWebsiteRequest ...
func TestPutBucketWebsiteRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutBucketWebsiteRequest(&s3.PutBucketWebsiteInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutObjectAclRequest ...
func TestPutObjectAclRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutObjectAclRequest(&s3.PutObjectAclInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutObjectLegalHoldRequest ...
func TestPutObjectLegalHoldRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutObjectLegalHoldRequest(&s3.PutObjectLegalHoldInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutObjectLockConfigurationRequest ...
func TestPutObjectLockConfigurationRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutObjectLockConfigurationRequest(&s3.PutObjectLockConfigurationInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutObjectRetentionRequest ...
func TestPutObjectRetentionRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutObjectRetentionRequest(&s3.PutObjectRetentionInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutObjectTaggingRequest ...
func TestPutObjectTaggingRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutObjectTaggingRequest(&s3.PutObjectTaggingInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//PutPublicAccessBlockRequest ...
func TestPutPublicAccessBlockRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.PutPublicAccessBlockRequest(&s3.PutPublicAccessBlockInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//RestoreObjectRequest ...
func TestRestoreObjectRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.RestoreObjectRequest(&s3.RestoreObjectInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//UploadPartRequest ...
func TestUploadPartRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.UploadPartRequest(&s3.UploadPartInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//UploadPartCopyRequest ...
func TestUploadPartCopyRequest(t *testing.T) {
	client := New(S3MemTestConfig)
	req := client.UploadPartCopyRequest(&s3.UploadPartCopyInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}

//WaitUntilBucketExists ...
func TestWaitUntilBucketExists(t *testing.T) {
	client := New(S3MemTestConfig)
	err := client.WaitUntilBucketExists(context.TODO(), &s3.HeadBucketInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, err.(s3memerr.S3MemError).Code())
}

//WaitUntilBucketNotExists ...
func TestWaitUntilBucketNotExists(t *testing.T) {
	client := New(S3MemTestConfig)
	err := client.WaitUntilBucketNotExists(context.TODO(), &s3.HeadBucketInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, err.(s3memerr.S3MemError).Code())
}

//WaitUntilObjectExists ...
func TestWaitUntilObjectExists(t *testing.T) {
	client := New(S3MemTestConfig)
	err := client.WaitUntilObjectExists(context.TODO(), &s3.HeadObjectInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, err.(s3memerr.S3MemError).Code())
}

//WaitUntilObjectNotExists ...
func TestWaitUntilObjectNotExists(t *testing.T) {
	client := New(S3MemTestConfig)
	err := client.WaitUntilObjectNotExists(context.TODO(), &s3.HeadObjectInput{})
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, err.(s3memerr.S3MemError).Code())
}
