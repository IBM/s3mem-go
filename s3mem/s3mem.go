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
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

const S3MemURL = "s3mem"

type Client struct {
	*aws.Client
}

var S3MemBuckets Buckets
var checkConfigHandler = aws.NamedHandler{Name: "s3mem.checkConfig", Fn: checkConfig}

func init() {
	S3MemBuckets.Buckets = make(map[string]*Bucket, 0)
}

func New(config aws.Config) *Client {
	endpoint, err := config.EndpointResolver.ResolveEndpoint(s3.EndpointsID, config.Region)
	if err != nil {
		endpoint = aws.Endpoint{}
	}
	svc := &Client{
		Client: &aws.Client{
			Metadata: aws.Metadata{
				ServiceName:   s3.ServiceName,
				ServiceID:     s3.ServiceID,
				SigningName:   endpoint.SigningName,
				SigningRegion: endpoint.SigningRegion,
				Endpoint:      endpoint.URL,
				APIVersion:    "2019-06-10",
			},
			Config:      config,
			Region:      config.Region,
			Credentials: config.Credentials,
			Handlers:    config.Handlers,
			Retryer:     config.Retryer,
			LogLevel:    config.LogLevel,
			Logger:      config.Logger,
		},
	}
	svc.Handlers.Build.PushBackNamed(checkConfigHandler)
	return svc
}

func checkConfig(r *aws.Request) {
	endpoint, err := r.Config.EndpointResolver.ResolveEndpoint(s3.EndpointsID, r.Config.Region)
	if err != nil {
		r.Error = s3memerr.NewError(err.Error(), "", nil, nil, nil, nil)
	}
	if endpoint.URL != S3MemURL {
		r.Error = s3memerr.NewError(s3memerr.ErrCodeNotS3MemRequest, "", nil, nil, nil, nil)
	}
}

func (c *Client) NotImplemented() *aws.Request {
	req := &aws.Request{
		HTTPRequest: &http.Request{},
	}
	req.Error = s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
	return req
}

//AbortMultipartUploadRequest ...
func (c *Client) AbortMultipartUploadRequest(input *s3.AbortMultipartUploadInput) s3.AbortMultipartUploadRequest {
	req := c.NotImplemented()
	return s3.AbortMultipartUploadRequest{Request: req, Input: input, Copy: nil}
}

//CompleteMultipartUploadRequest ...
func (c *Client) CompleteMultipartUploadRequest(input *s3.CompleteMultipartUploadInput) s3.CompleteMultipartUploadRequest {
	req := c.NotImplemented()
	return s3.CompleteMultipartUploadRequest{Request: req, Input: input, Copy: nil}
}

//CopyObjectRequest ...
func (c *Client) CopyObjectRequest(input *s3.CopyObjectInput) s3.CopyObjectRequest {
	req := c.NotImplemented()
	return s3.CopyObjectRequest{Request: req, Input: input, Copy: nil}
}

//CreateMultipartUploadRequest ...
func (c *Client) CreateMultipartUploadRequest(input *s3.CreateMultipartUploadInput) s3.CreateMultipartUploadRequest {
	req := c.NotImplemented()
	return s3.CreateMultipartUploadRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketAnalyticsConfigurationRequest ...
func (c *Client) DeleteBucketAnalyticsConfigurationRequest(input *s3.DeleteBucketAnalyticsConfigurationInput) s3.DeleteBucketAnalyticsConfigurationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketCorsRequest ...
func (c *Client) DeleteBucketCorsRequest(input *s3.DeleteBucketCorsInput) s3.DeleteBucketCorsRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketCorsRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketEncryptionRequest ...
func (c *Client) DeleteBucketEncryptionRequest(input *s3.DeleteBucketEncryptionInput) s3.DeleteBucketEncryptionRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketEncryptionRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketInventoryConfigurationRequest ...
func (c *Client) DeleteBucketInventoryConfigurationRequest(input *s3.DeleteBucketInventoryConfigurationInput) s3.DeleteBucketInventoryConfigurationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketLifecycleRequest ...
func (c *Client) DeleteBucketLifecycleRequest(input *s3.DeleteBucketLifecycleInput) s3.DeleteBucketLifecycleRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketLifecycleRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketMetricsConfigurationRequest ...
func (c *Client) DeleteBucketMetricsConfigurationRequest(input *s3.DeleteBucketMetricsConfigurationInput) s3.DeleteBucketMetricsConfigurationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketPolicyRequest ...
func (c *Client) DeleteBucketPolicyRequest(input *s3.DeleteBucketPolicyInput) s3.DeleteBucketPolicyRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketPolicyRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketReplicationRequest ...
func (c *Client) DeleteBucketReplicationRequest(input *s3.DeleteBucketReplicationInput) s3.DeleteBucketReplicationRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketReplicationRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketTaggingRequest ...
func (c *Client) DeleteBucketTaggingRequest(input *s3.DeleteBucketTaggingInput) s3.DeleteBucketTaggingRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketTaggingRequest{Request: req, Input: input, Copy: nil}
}

//DeleteBucketWebsiteRequest ...
func (c *Client) DeleteBucketWebsiteRequest(input *s3.DeleteBucketWebsiteInput) s3.DeleteBucketWebsiteRequest {
	req := c.NotImplemented()
	return s3.DeleteBucketWebsiteRequest{Request: req, Input: input, Copy: nil}
}

//DeleteObjectTaggingRequest ...
func (c *Client) DeleteObjectTaggingRequest(input *s3.DeleteObjectTaggingInput) s3.DeleteObjectTaggingRequest {
	req := c.NotImplemented()
	return s3.DeleteObjectTaggingRequest{Request: req, Input: input, Copy: nil}
}

//DeletePublicAccessBlockRequest ...
func (c *Client) DeletePublicAccessBlockRequest(input *s3.DeletePublicAccessBlockInput) s3.DeletePublicAccessBlockRequest {
	req := c.NotImplemented()
	return s3.DeletePublicAccessBlockRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketAccelerateConfigurationRequest ...
func (c *Client) GetBucketAccelerateConfigurationRequest(input *s3.GetBucketAccelerateConfigurationInput) s3.GetBucketAccelerateConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketAccelerateConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketAclRequest ...
func (c *Client) GetBucketAclRequest(input *s3.GetBucketAclInput) s3.GetBucketAclRequest {
	req := c.NotImplemented()
	return s3.GetBucketAclRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketAnalyticsConfigurationRequest ...
func (c *Client) GetBucketAnalyticsConfigurationRequest(input *s3.GetBucketAnalyticsConfigurationInput) s3.GetBucketAnalyticsConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketCorsRequest ...
func (c *Client) GetBucketCorsRequest(input *s3.GetBucketCorsInput) s3.GetBucketCorsRequest {
	req := c.NotImplemented()
	return s3.GetBucketCorsRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketEncryptionRequest ...
func (c *Client) GetBucketEncryptionRequest(input *s3.GetBucketEncryptionInput) s3.GetBucketEncryptionRequest {
	req := c.NotImplemented()
	return s3.GetBucketEncryptionRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketInventoryConfigurationRequest ...
func (c *Client) GetBucketInventoryConfigurationRequest(input *s3.GetBucketInventoryConfigurationInput) s3.GetBucketInventoryConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketLifecycleRequest ...
func (c *Client) GetBucketLifecycleRequest(input *s3.GetBucketLifecycleInput) s3.GetBucketLifecycleRequest {
	req := c.NotImplemented()
	return s3.GetBucketLifecycleRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketLifecycleConfigurationRequest ...
func (c *Client) GetBucketLifecycleConfigurationRequest(input *s3.GetBucketLifecycleConfigurationInput) s3.GetBucketLifecycleConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketLifecycleConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketLocationRequest ...
func (c *Client) GetBucketLocationRequest(input *s3.GetBucketLocationInput) s3.GetBucketLocationRequest {
	req := c.NotImplemented()
	return s3.GetBucketLocationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketLoggingRequest ...
func (c *Client) GetBucketLoggingRequest(input *s3.GetBucketLoggingInput) s3.GetBucketLoggingRequest {
	req := c.NotImplemented()
	return s3.GetBucketLoggingRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketMetricsConfigurationRequest ...
func (c *Client) GetBucketMetricsConfigurationRequest(input *s3.GetBucketMetricsConfigurationInput) s3.GetBucketMetricsConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketNotificationRequest ...
func (c *Client) GetBucketNotificationRequest(input *s3.GetBucketNotificationInput) s3.GetBucketNotificationRequest {
	req := c.NotImplemented()
	return s3.GetBucketNotificationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketNotificationConfigurationRequest ...
func (c *Client) GetBucketNotificationConfigurationRequest(input *s3.GetBucketNotificationConfigurationInput) s3.GetBucketNotificationConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetBucketNotificationConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketPolicyRequest ...
func (c *Client) GetBucketPolicyRequest(input *s3.GetBucketPolicyInput) s3.GetBucketPolicyRequest {
	req := c.NotImplemented()
	return s3.GetBucketPolicyRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketPolicyStatusRequest ...
func (c *Client) GetBucketPolicyStatusRequest(input *s3.GetBucketPolicyStatusInput) s3.GetBucketPolicyStatusRequest {
	req := c.NotImplemented()
	return s3.GetBucketPolicyStatusRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketReplicationRequest ...
func (c *Client) GetBucketReplicationRequest(input *s3.GetBucketReplicationInput) s3.GetBucketReplicationRequest {
	req := c.NotImplemented()
	return s3.GetBucketReplicationRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketRequestPaymentRequest ...
func (c *Client) GetBucketRequestPaymentRequest(input *s3.GetBucketRequestPaymentInput) s3.GetBucketRequestPaymentRequest {
	req := c.NotImplemented()
	return s3.GetBucketRequestPaymentRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketTaggingRequest ...
func (c *Client) GetBucketTaggingRequest(input *s3.GetBucketTaggingInput) s3.GetBucketTaggingRequest {
	req := c.NotImplemented()
	return s3.GetBucketTaggingRequest{Request: req, Input: input, Copy: nil}
}

//GetBucketWebsiteRequest ...
func (c *Client) GetBucketWebsiteRequest(input *s3.GetBucketWebsiteInput) s3.GetBucketWebsiteRequest {
	req := c.NotImplemented()
	return s3.GetBucketWebsiteRequest{Request: req, Input: input, Copy: nil}
}

//GetObjectAclRequest ...
func (c *Client) GetObjectAclRequest(input *s3.GetObjectAclInput) s3.GetObjectAclRequest {
	req := c.NotImplemented()
	return s3.GetObjectAclRequest{Request: req, Input: input, Copy: nil}
}

//GetObjectLegalHoldRequest ...
func (c *Client) GetObjectLegalHoldRequest(input *s3.GetObjectLegalHoldInput) s3.GetObjectLegalHoldRequest {
	req := c.NotImplemented()
	return s3.GetObjectLegalHoldRequest{Request: req, Input: input, Copy: nil}
}

//GetObjectLockConfigurationRequest ...
func (c *Client) GetObjectLockConfigurationRequest(input *s3.GetObjectLockConfigurationInput) s3.GetObjectLockConfigurationRequest {
	req := c.NotImplemented()
	return s3.GetObjectLockConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//GetObjectRetentionRequest ...
func (c *Client) GetObjectRetentionRequest(input *s3.GetObjectRetentionInput) s3.GetObjectRetentionRequest {
	req := c.NotImplemented()
	return s3.GetObjectRetentionRequest{Request: req, Input: input, Copy: nil}
}

//GetObjectTaggingRequest ...
func (c *Client) GetObjectTaggingRequest(input *s3.GetObjectTaggingInput) s3.GetObjectTaggingRequest {
	req := c.NotImplemented()
	return s3.GetObjectTaggingRequest{Request: req, Input: input, Copy: nil}
}

//GetObjectTorrentRequest ...
func (c *Client) GetObjectTorrentRequest(input *s3.GetObjectTorrentInput) s3.GetObjectTorrentRequest {
	req := c.NotImplemented()
	return s3.GetObjectTorrentRequest{Request: req, Input: input, Copy: nil}
}

//GetPublicAccessBlockRequest ...
func (c *Client) GetPublicAccessBlockRequest(input *s3.GetPublicAccessBlockInput) s3.GetPublicAccessBlockRequest {
	req := c.NotImplemented()
	return s3.GetPublicAccessBlockRequest{Request: req, Input: input, Copy: nil}
}

//HeadBucketRequest ...
func (c *Client) HeadBucketRequest(input *s3.HeadBucketInput) s3.HeadBucketRequest {
	req := c.NotImplemented()
	return s3.HeadBucketRequest{Request: req, Input: input, Copy: nil}
}

//HeadObjectRequest ...
func (c *Client) HeadObjectRequest(input *s3.HeadObjectInput) s3.HeadObjectRequest {
	req := c.NotImplemented()
	return s3.HeadObjectRequest{Request: req, Input: input, Copy: nil}
}

//ListBucketAnalyticsConfigurationsRequest ...
func (c *Client) ListBucketAnalyticsConfigurationsRequest(input *s3.ListBucketAnalyticsConfigurationsInput) s3.ListBucketAnalyticsConfigurationsRequest {
	req := c.NotImplemented()
	return s3.ListBucketAnalyticsConfigurationsRequest{Request: req, Input: input, Copy: nil}
}

//ListBucketInventoryConfigurationsRequest ...
func (c *Client) ListBucketInventoryConfigurationsRequest(input *s3.ListBucketInventoryConfigurationsInput) s3.ListBucketInventoryConfigurationsRequest {
	req := c.NotImplemented()
	return s3.ListBucketInventoryConfigurationsRequest{Request: req, Input: input, Copy: nil}
}

//ListBucketMetricsConfigurationsRequest ...
func (c *Client) ListBucketMetricsConfigurationsRequest(input *s3.ListBucketMetricsConfigurationsInput) s3.ListBucketMetricsConfigurationsRequest {
	req := c.NotImplemented()
	return s3.ListBucketMetricsConfigurationsRequest{Request: req, Input: input, Copy: nil}
}

//ListMultipartUploadsRequest ...
func (c *Client) ListMultipartUploadsRequest(input *s3.ListMultipartUploadsInput) s3.ListMultipartUploadsRequest {
	req := c.NotImplemented()
	return s3.ListMultipartUploadsRequest{Request: req, Input: input, Copy: nil}
}

//ListObjectVersionsRequest ...
func (c *Client) ListObjectVersionsRequest(input *s3.ListObjectVersionsInput) s3.ListObjectVersionsRequest {
	req := c.NotImplemented()
	return s3.ListObjectVersionsRequest{Request: req, Input: input, Copy: nil}
}

//ListObjectsV2Request ...
func (c *Client) ListObjectsV2Request(input *s3.ListObjectsV2Input) s3.ListObjectsV2Request {
	req := c.NotImplemented()
	return s3.ListObjectsV2Request{Request: req, Input: input, Copy: nil}
}

//ListPartsRequest ...
func (c *Client) ListPartsRequest(input *s3.ListPartsInput) s3.ListPartsRequest {
	req := c.NotImplemented()
	return s3.ListPartsRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketAccelerateConfigurationRequest ...
func (c *Client) PutBucketAccelerateConfigurationRequest(input *s3.PutBucketAccelerateConfigurationInput) s3.PutBucketAccelerateConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketAccelerateConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketAclRequest ...
func (c *Client) PutBucketAclRequest(input *s3.PutBucketAclInput) s3.PutBucketAclRequest {
	req := c.NotImplemented()
	return s3.PutBucketAclRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketAnalyticsConfigurationRequest ...
func (c *Client) PutBucketAnalyticsConfigurationRequest(input *s3.PutBucketAnalyticsConfigurationInput) s3.PutBucketAnalyticsConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketAnalyticsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketCorsRequest ...
func (c *Client) PutBucketCorsRequest(input *s3.PutBucketCorsInput) s3.PutBucketCorsRequest {
	req := c.NotImplemented()
	return s3.PutBucketCorsRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketEncryptionRequest ...
func (c *Client) PutBucketEncryptionRequest(input *s3.PutBucketEncryptionInput) s3.PutBucketEncryptionRequest {
	req := c.NotImplemented()
	return s3.PutBucketEncryptionRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketInventoryConfigurationRequest ...
func (c *Client) PutBucketInventoryConfigurationRequest(input *s3.PutBucketInventoryConfigurationInput) s3.PutBucketInventoryConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketInventoryConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketLifecycleRequest ...
func (c *Client) PutBucketLifecycleRequest(input *s3.PutBucketLifecycleInput) s3.PutBucketLifecycleRequest {
	req := c.NotImplemented()
	return s3.PutBucketLifecycleRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketLifecycleConfigurationRequest ...
func (c *Client) PutBucketLifecycleConfigurationRequest(input *s3.PutBucketLifecycleConfigurationInput) s3.PutBucketLifecycleConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketLifecycleConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketLoggingRequest ...
func (c *Client) PutBucketLoggingRequest(input *s3.PutBucketLoggingInput) s3.PutBucketLoggingRequest {
	req := c.NotImplemented()
	return s3.PutBucketLoggingRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketMetricsConfigurationRequest ...
func (c *Client) PutBucketMetricsConfigurationRequest(input *s3.PutBucketMetricsConfigurationInput) s3.PutBucketMetricsConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketMetricsConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketNotificationRequest ...
func (c *Client) PutBucketNotificationRequest(input *s3.PutBucketNotificationInput) s3.PutBucketNotificationRequest {
	req := c.NotImplemented()
	return s3.PutBucketNotificationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketNotificationConfigurationRequest ...
func (c *Client) PutBucketNotificationConfigurationRequest(input *s3.PutBucketNotificationConfigurationInput) s3.PutBucketNotificationConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutBucketNotificationConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketPolicyRequest ...
func (c *Client) PutBucketPolicyRequest(input *s3.PutBucketPolicyInput) s3.PutBucketPolicyRequest {
	req := c.NotImplemented()
	return s3.PutBucketPolicyRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketReplicationRequest ...
func (c *Client) PutBucketReplicationRequest(input *s3.PutBucketReplicationInput) s3.PutBucketReplicationRequest {
	req := c.NotImplemented()
	return s3.PutBucketReplicationRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketRequestPaymentRequest ...
func (c *Client) PutBucketRequestPaymentRequest(input *s3.PutBucketRequestPaymentInput) s3.PutBucketRequestPaymentRequest {
	req := c.NotImplemented()
	return s3.PutBucketRequestPaymentRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketTaggingRequest ...
func (c *Client) PutBucketTaggingRequest(input *s3.PutBucketTaggingInput) s3.PutBucketTaggingRequest {
	req := c.NotImplemented()
	return s3.PutBucketTaggingRequest{Request: req, Input: input, Copy: nil}
}

//PutBucketWebsiteRequest ...
func (c *Client) PutBucketWebsiteRequest(input *s3.PutBucketWebsiteInput) s3.PutBucketWebsiteRequest {
	req := c.NotImplemented()
	return s3.PutBucketWebsiteRequest{Request: req, Input: input, Copy: nil}
}

//PutObjectAclRequest ...
func (c *Client) PutObjectAclRequest(input *s3.PutObjectAclInput) s3.PutObjectAclRequest {
	req := c.NotImplemented()
	return s3.PutObjectAclRequest{Request: req, Input: input, Copy: nil}
}

//PutObjectLegalHoldRequest ...
func (c *Client) PutObjectLegalHoldRequest(input *s3.PutObjectLegalHoldInput) s3.PutObjectLegalHoldRequest {
	req := c.NotImplemented()
	return s3.PutObjectLegalHoldRequest{Request: req, Input: input, Copy: nil}
}

//PutObjectLockConfigurationRequest ...
func (c *Client) PutObjectLockConfigurationRequest(input *s3.PutObjectLockConfigurationInput) s3.PutObjectLockConfigurationRequest {
	req := c.NotImplemented()
	return s3.PutObjectLockConfigurationRequest{Request: req, Input: input, Copy: nil}
}

//PutObjectRetentionRequest ...
func (c *Client) PutObjectRetentionRequest(input *s3.PutObjectRetentionInput) s3.PutObjectRetentionRequest {
	req := c.NotImplemented()
	return s3.PutObjectRetentionRequest{Request: req, Input: input, Copy: nil}
}

//PutObjectTaggingRequest ...
func (c *Client) PutObjectTaggingRequest(input *s3.PutObjectTaggingInput) s3.PutObjectTaggingRequest {
	req := c.NotImplemented()
	return s3.PutObjectTaggingRequest{Request: req, Input: input, Copy: nil}
}

//PutPublicAccessBlockRequest ...
func (c *Client) PutPublicAccessBlockRequest(input *s3.PutPublicAccessBlockInput) s3.PutPublicAccessBlockRequest {
	req := c.NotImplemented()
	return s3.PutPublicAccessBlockRequest{Request: req, Input: input, Copy: nil}
}

//RestoreObjectRequest ...
func (c *Client) RestoreObjectRequest(input *s3.RestoreObjectInput) s3.RestoreObjectRequest {
	req := c.NotImplemented()
	return s3.RestoreObjectRequest{Request: req, Input: input, Copy: nil}
}

//UploadPartRequest ...
func (c *Client) UploadPartRequest(input *s3.UploadPartInput) s3.UploadPartRequest {
	req := c.NotImplemented()
	return s3.UploadPartRequest{Request: req, Input: input, Copy: nil}
}

//UploadPartCopyRequest ...
func (c *Client) UploadPartCopyRequest(input *s3.UploadPartCopyInput) s3.UploadPartCopyRequest {
	req := c.NotImplemented()
	return s3.UploadPartCopyRequest{Request: req, Input: input, Copy: nil}
}

//WaitUntilBucketExists ...
func (c *Client) WaitUntilBucketExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}

//WaitUntilBucketNotExists ...
func (c *Client) WaitUntilBucketNotExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}

//WaitUntilObjectExists ...
func (c *Client) WaitUntilObjectExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}

//WaitUntilObjectNotExists ...
func (c *Client) WaitUntilObjectNotExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error {
	return s3memerr.NewError(s3memerr.ErrCodeNotImplemented, "Not implemented", nil, nil, nil, nil)
}
