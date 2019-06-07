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
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

const S3MemURL = "s3mem"

type S3Mem struct {
	*aws.Client
}

var S3MemBuckets Buckets

func init() {
	S3MemBuckets.Buckets = make(map[string]*Bucket, 0)
}

func New(config aws.Config) s3iface.S3API {
	var signingName string
	signingRegion := config.Region

	svc := &S3Mem{
		Client: aws.NewClient(
			config,
			aws.Metadata{
				ServiceName:   s3.ServiceName,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				APIVersion:    "2006-03-01",
			},
		),
	}
	return svc
}

func (c *S3Mem) NotImplemented() *aws.Request {
	req := &aws.Request{
		HTTPRequest: &http.Request{},
	}
	req.Error = s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
	return req
}

func (c *S3Mem) AbortMultipartUploadRequest(input *s3.AbortMultipartUploadInput) s3.AbortMultipartUploadRequest {
	req := c.NotImplemented()
	return s3.AbortMultipartUploadRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) CompleteMultipartUploadRequest(input *s3.CompleteMultipartUploadInput) s3.CompleteMultipartUploadRequest {
	req := c.NotImplemented()
	return s3.CompleteMultipartUploadRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) CopyObjectRequest(input *s3.CopyObjectInput) s3.CopyObjectRequest {
	req := c.NotImplemented()
	return s3.CopyObjectRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) CreateMultipartUploadRequest(input *s3.CreateMultipartUploadInput) s3.CreateMultipartUploadRequest {
	req := c.NotImplemented()
	return s3.CreateMultipartUploadRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketAnalyticsConfigurationRequest(input *s3.DeleteBucketAnalyticsConfigurationInput) s3.DeleteBucketAnalyticsConfigurationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketCorsRequest(input *s3.DeleteBucketCorsInput) s3.DeleteBucketCorsRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketCorsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketEncryptionRequest(input *s3.DeleteBucketEncryptionInput) s3.DeleteBucketEncryptionRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketEncryptionRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketInventoryConfigurationRequest(input *s3.DeleteBucketInventoryConfigurationInput) s3.DeleteBucketInventoryConfigurationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketLifecycleRequest(input *s3.DeleteBucketLifecycleInput) s3.DeleteBucketLifecycleRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketLifecycleRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketMetricsConfigurationRequest(input *s3.DeleteBucketMetricsConfigurationInput) s3.DeleteBucketMetricsConfigurationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketPolicyRequest(input *s3.DeleteBucketPolicyInput) s3.DeleteBucketPolicyRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketPolicyRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketReplicationRequest(input *s3.DeleteBucketReplicationInput) s3.DeleteBucketReplicationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketReplicationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketTaggingRequest(input *s3.DeleteBucketTaggingInput) s3.DeleteBucketTaggingRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketTaggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteBucketWebsiteRequest(input *s3.DeleteBucketWebsiteInput) s3.DeleteBucketWebsiteRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketWebsiteRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeleteObjectTaggingRequest(input *s3.DeleteObjectTaggingInput) s3.DeleteObjectTaggingRequest {
	req := c.NotImplemented()
	return s3.DeleteObjectTaggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) DeletePublicAccessBlockRequest(input *s3.DeletePublicAccessBlockInput) s3.DeletePublicAccessBlockRequest {
	req := c.NotImplemented()
	return s3.DeletePublicAccessBlockRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketAccelerateConfigurationRequest(input *s3.GetBucketAccelerateConfigurationInput) s3.GetBucketAccelerateConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketAccelerateConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketAclRequest(input *s3.GetBucketAclInput) s3.GetBucketAclRequest {
	req := c.NotImplemented()
	return s3.GetBucketAclRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketAnalyticsConfigurationRequest(input *s3.GetBucketAnalyticsConfigurationInput) s3.GetBucketAnalyticsConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketCorsRequest(input *s3.GetBucketCorsInput) s3.GetBucketCorsRequest {
	req := c.NotImplemented()
	return s3.GetBucketCorsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketEncryptionRequest(input *s3.GetBucketEncryptionInput) s3.GetBucketEncryptionRequest {
	req := c.NotImplemented()
	return s3.GetBucketEncryptionRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketInventoryConfigurationRequest(input *s3.GetBucketInventoryConfigurationInput) s3.GetBucketInventoryConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketLifecycleRequest(input *s3.GetBucketLifecycleInput) s3.GetBucketLifecycleRequest {
	req := c.NotImplemented()
	return s3.GetBucketLifecycleRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketLifecycleConfigurationRequest(input *s3.GetBucketLifecycleConfigurationInput) s3.GetBucketLifecycleConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketLifecycleConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketLocationRequest(input *s3.GetBucketLocationInput) s3.GetBucketLocationRequest {
	req := c.NotImplemented()
	return s3.GetBucketLocationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketLoggingRequest(input *s3.GetBucketLoggingInput) s3.GetBucketLoggingRequest {
	req := c.NotImplemented()
	return s3.GetBucketLoggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketMetricsConfigurationRequest(input *s3.GetBucketMetricsConfigurationInput) s3.GetBucketMetricsConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketNotificationRequest(input *s3.GetBucketNotificationConfigurationInput) s3.GetBucketNotificationRequest {
	req := c.NotImplemented()
	return s3.GetBucketNotificationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketNotificationConfigurationRequest(input *s3.GetBucketNotificationConfigurationInput) s3.GetBucketNotificationConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketNotificationConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketPolicyRequest(input *s3.GetBucketPolicyInput) s3.GetBucketPolicyRequest {
	req := c.NotImplemented()
	return s3.GetBucketPolicyRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketPolicyStatusRequest(input *s3.GetBucketPolicyStatusInput) s3.GetBucketPolicyStatusRequest {
	req := c.NotImplemented()
	return s3.GetBucketPolicyStatusRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketReplicationRequest(input *s3.GetBucketReplicationInput) s3.GetBucketReplicationRequest {
	req := c.NotImplemented()
	return s3.GetBucketReplicationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketRequestPaymentRequest(input *s3.GetBucketRequestPaymentInput) s3.GetBucketRequestPaymentRequest {
	req := c.NotImplemented()
	return s3.GetBucketRequestPaymentRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketTaggingRequest(input *s3.GetBucketTaggingInput) s3.GetBucketTaggingRequest {
	req := c.NotImplemented()
	return s3.GetBucketTaggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetBucketWebsiteRequest(input *s3.GetBucketWebsiteInput) s3.GetBucketWebsiteRequest {
	req := c.NotImplemented()
	return s3.GetBucketWebsiteRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetObjectAclRequest(input *s3.GetObjectAclInput) s3.GetObjectAclRequest {
	req := c.NotImplemented()
	return s3.GetObjectAclRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetObjectLegalHoldRequest(input *s3.GetObjectLegalHoldInput) s3.GetObjectLegalHoldRequest {
	req := c.NotImplemented()
	return s3.GetObjectLegalHoldRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetObjectLockConfigurationRequest(input *s3.GetObjectLockConfigurationInput) s3.GetObjectLockConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetObjectLockConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetObjectRetentionRequest(input *s3.GetObjectRetentionInput) s3.GetObjectRetentionRequest {
	req := c.NotImplemented()
	return s3.GetObjectRetentionRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetObjectTaggingRequest(input *s3.GetObjectTaggingInput) s3.GetObjectTaggingRequest {
	req := c.NotImplemented()
	return s3.GetObjectTaggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetObjectTorrentRequest(input *s3.GetObjectTorrentInput) s3.GetObjectTorrentRequest {
	req := c.NotImplemented()
	return s3.GetObjectTorrentRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) GetPublicAccessBlockRequest(input *s3.GetPublicAccessBlockInput) s3.GetPublicAccessBlockRequest {
	req := c.NotImplemented()
	return s3.GetPublicAccessBlockRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) HeadBucketRequest(input *s3.HeadBucketInput) s3.HeadBucketRequest {
	req := c.NotImplemented()
	return s3.HeadBucketRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) HeadObjectRequest(input *s3.HeadObjectInput) s3.HeadObjectRequest {
	req := c.NotImplemented()
	return s3.HeadObjectRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) ListBucketAnalyticsConfigurationsRequest(input *s3.ListBucketAnalyticsConfigurationsInput) s3.ListBucketAnalyticsConfigurationsRequest {
	req := c.NotImplemented()
	return s3.ListBucketAnalyticsConfigurationsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) ListBucketInventoryConfigurationsRequest(input *s3.ListBucketInventoryConfigurationsInput) s3.ListBucketInventoryConfigurationsRequest {
	req := c.NotImplemented()
	return s3.ListBucketInventoryConfigurationsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) ListBucketMetricsConfigurationsRequest(input *s3.ListBucketMetricsConfigurationsInput) s3.ListBucketMetricsConfigurationsRequest {
	req := c.NotImplemented()
	return s3.ListBucketMetricsConfigurationsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) ListMultipartUploadsRequest(input *s3.ListMultipartUploadsInput) s3.ListMultipartUploadsRequest {
	req := c.NotImplemented()
	return s3.ListMultipartUploadsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) ListObjectVersionsRequest(input *s3.ListObjectVersionsInput) s3.ListObjectVersionsRequest {
	req := c.NotImplemented()
	return s3.ListObjectVersionsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) ListObjectsV2Request(input *s3.ListObjectsV2Input) s3.ListObjectsV2Request {
	req := c.NotImplemented()
	return s3.ListObjectsV2Request{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) ListPartsRequest(input *s3.ListPartsInput) s3.ListPartsRequest {
	req := c.NotImplemented()
	return s3.ListPartsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketAccelerateConfigurationRequest(input *s3.PutBucketAccelerateConfigurationInput) s3.PutBucketAccelerateConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketAccelerateConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketAclRequest(input *s3.PutBucketAclInput) s3.PutBucketAclRequest {
	req := c.NotImplemented()
	return s3.PutBucketAclRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketAnalyticsConfigurationRequest(input *s3.PutBucketAnalyticsConfigurationInput) s3.PutBucketAnalyticsConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketCorsRequest(input *s3.PutBucketCorsInput) s3.PutBucketCorsRequest {
	req := c.NotImplemented()
	return s3.PutBucketCorsRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketEncryptionRequest(input *s3.PutBucketEncryptionInput) s3.PutBucketEncryptionRequest {
	req := c.NotImplemented()
	return s3.PutBucketEncryptionRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketInventoryConfigurationRequest(input *s3.PutBucketInventoryConfigurationInput) s3.PutBucketInventoryConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketLifecycleRequest(input *s3.PutBucketLifecycleInput) s3.PutBucketLifecycleRequest {
	req := c.NotImplemented()
	return s3.PutBucketLifecycleRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketLifecycleConfigurationRequest(input *s3.PutBucketLifecycleConfigurationInput) s3.PutBucketLifecycleConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketLifecycleConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketLoggingRequest(input *s3.PutBucketLoggingInput) s3.PutBucketLoggingRequest {
	req := c.NotImplemented()
	return s3.PutBucketLoggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketMetricsConfigurationRequest(input *s3.PutBucketMetricsConfigurationInput) s3.PutBucketMetricsConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketNotificationRequest(input *s3.PutBucketNotificationInput) s3.PutBucketNotificationRequest {
	req := c.NotImplemented()
	return s3.PutBucketNotificationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketNotificationConfigurationRequest(input *s3.PutBucketNotificationConfigurationInput) s3.PutBucketNotificationConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketNotificationConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketPolicyRequest(input *s3.PutBucketPolicyInput) s3.PutBucketPolicyRequest {
	req := c.NotImplemented()
	return s3.PutBucketPolicyRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketReplicationRequest(input *s3.PutBucketReplicationInput) s3.PutBucketReplicationRequest {
	req := c.NotImplemented()
	return s3.PutBucketReplicationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketRequestPaymentRequest(input *s3.PutBucketRequestPaymentInput) s3.PutBucketRequestPaymentRequest {
	req := c.NotImplemented()
	return s3.PutBucketRequestPaymentRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketTaggingRequest(input *s3.PutBucketTaggingInput) s3.PutBucketTaggingRequest {
	req := c.NotImplemented()
	return s3.PutBucketTaggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutBucketWebsiteRequest(input *s3.PutBucketWebsiteInput) s3.PutBucketWebsiteRequest {
	req := c.NotImplemented()
	return s3.PutBucketWebsiteRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutObjectAclRequest(input *s3.PutObjectAclInput) s3.PutObjectAclRequest {
	req := c.NotImplemented()
	return s3.PutObjectAclRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutObjectLegalHoldRequest(input *s3.PutObjectLegalHoldInput) s3.PutObjectLegalHoldRequest {
	req := c.NotImplemented()
	return s3.PutObjectLegalHoldRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutObjectLockConfigurationRequest(input *s3.PutObjectLockConfigurationInput) s3.PutObjectLockConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutObjectLockConfigurationRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutObjectRetentionRequest(input *s3.PutObjectRetentionInput) s3.PutObjectRetentionRequest {
	req := c.NotImplemented()
	return s3.PutObjectRetentionRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutObjectTaggingRequest(input *s3.PutObjectTaggingInput) s3.PutObjectTaggingRequest {
	req := c.NotImplemented()
	return s3.PutObjectTaggingRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) PutPublicAccessBlockRequest(input *s3.PutPublicAccessBlockInput) s3.PutPublicAccessBlockRequest {
	req := c.NotImplemented()
	return s3.PutPublicAccessBlockRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) RestoreObjectRequest(input *s3.RestoreObjectInput) s3.RestoreObjectRequest {
	req := c.NotImplemented()
	return s3.RestoreObjectRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) UploadPartRequest(input *s3.UploadPartInput) s3.UploadPartRequest {
	req := c.NotImplemented()
	return s3.UploadPartRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) UploadPartCopyRequest(input *s3.UploadPartCopyInput) s3.UploadPartCopyRequest {
	req := c.NotImplemented()
	return s3.UploadPartCopyRequest{Request: req, Input: input, Copy: nil}
}

func (c *S3Mem) WaitUntilBucketExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}

func (c *S3Mem) WaitUntilBucketNotExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}

func (c *S3Mem) WaitUntilObjectExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}

func (c *S3Mem) WaitUntilObjectNotExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}
