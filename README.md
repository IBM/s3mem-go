# s3mem
s3mem is an in-memory S3API implementation. 
It is a real serverless AWS S3, no need to run an external process or a docker image. It is a Go package that you use, that's it!
It allows you to create different S3 servers for your tests, for example one per unit-test.
It is a work in progess and only some APIs are implemented (Feel free to contribute)

## Build

This project uses [dep](https://github.com/golang/dep), running `dep ensure -v` will update the vendor directory with all dependencies.
If your projectg doesn't use `dep` then you will need to make available all packages of the vendor directory to your project.

## Usage

To get a client call the `s3mem.New()` function. It returns a s3iface.ClientAPI client for the s3mem in-memory implementation and you can use it as client for the AWS S3 methods.

### Examples

As example here a method which retrieves an object from an S3 implementation [Example](example/example.go) and here how to write a test for it [Example Test](example/example_test.go) using s3mem. It shows also how to manage multiple s3 servers.
The s3mem package offers a number of helper methods to manage the buckets and objects without using the S3API. These methods will directly modify the in-memory s3. This is useful to setup your tests. You can find them in the [Helper methods](s3mem/helper.go). To create new S3 server call `NewDefaultS3MemService()` or `NewS3MemService(<s3-name>)` and to get the store `GetDefaultS3MemService()` or `GetS3MemService(<s3-name>)`.

## Implemented methods

```go
CopyObjectRequest(*s3.CopyObjectInput) s3.CopyObjectRequest
CreateBucketRequest(input *s3.CreateBucketInput) s3.CreateBucketRequest
DeleteBucketRequest(input *s3.DeleteBucketInput) s3.DeleteBucketRequest
DeleteObjectRequest(input *s3.DeleteObjectInput) s3.DeleteObjectRequest
DeleteObjectsRequest(input *s3.DeleteObjectsInput) s3.DeleteObjectsRequest
GetBucketVersioningRequest(input *s3.GetBucketVersioningInput) s3.GetBucketVersioningRequest
GetObjectRequest(input *s3.GetObjectInput) s3.GetObjectRequest
ListBucketsRequest(input *s3.ListBucketsInput) s3.ListBucketsRequest
ListObjectsRequest(input *s3.ListObjectsInput) s3.ListObjectsRequest
PutBucketVersioningRequest(input *s3.PutBucketVersioningInput) s3.PutBucketVersioningRequest
PutObjectRequest(input *s3.PutObjectInput) s3.PutObjectRequest
```

## Not implemented methods

A call to Non-implemented methods will return a ErrCodeNotImplemented error.

```go
AbortMultipartUploadRequest(*s3.AbortMultipartUploadInput) s3.AbortMultipartUploadRequest
CompleteMultipartUploadRequest(*s3.CompleteMultipartUploadInput) s3.CompleteMultipartUploadRequest
CreateMultipartUploadRequest(*s3.CreateMultipartUploadInput) s3.CreateMultipartUploadRequest
DeleteBucketAnalyticsConfigurationRequest(*s3.DeleteBucketAnalyticsConfigurationInput) s3.DeleteBucketAnalyticsConfigurationRequest
DeleteBucketCorsRequest(*s3.DeleteBucketCorsInput) s3.DeleteBucketCorsRequest 
DeleteBucketEncryptionRequest(*s3.DeleteBucketEncryptionInput) s3.DeleteBucketEncryptionRequest
DeleteBucketInventoryConfigurationRequest(*s3.DeleteBucketInventoryConfigurationInput) s3.DeleteBucketInventoryConfigurationRequest 
DeleteBucketLifecycleRequest(*s3.DeleteBucketLifecycleInput) s3.DeleteBucketLifecycleRequest 
DeleteBucketMetricsConfigurationRequest(*s3.DeleteBucketMetricsConfigurationInput) s3.DeleteBucketMetricsConfigurationRequest 
DeleteBucketPolicyRequest(*s3.DeleteBucketPolicyInput) s3.DeleteBucketPolicyRequest 
DeleteBucketReplicationRequest(*s3.DeleteBucketReplicationInput) s3.DeleteBucketReplicationRequest 
DeleteBucketTaggingRequest(*s3.DeleteBucketTaggingInput) s3.DeleteBucketTaggingRequest 
DeleteBucketWebsiteRequest(*s3.DeleteBucketWebsiteInput) s3.DeleteBucketWebsiteRequest 
DeleteObjectTaggingRequest(*s3.DeleteObjectTaggingInput) s3.DeleteObjectTaggingRequest 
DeletePublicAccessBlockRequest(*s3.DeletePublicAccessBlockInput) s3.DeletePublicAccessBlockRequest 
GetBucketAccelerateConfigurationRequest(*s3.GetBucketAccelerateConfigurationInput) s3.GetBucketAccelerateConfigurationRequest 
GetBucketAclRequest(*s3.GetBucketAclInput) s3.GetBucketAclRequest 
GetBucketAnalyticsConfigurationRequest(*s3.GetBucketAnalyticsConfigurationInput) s3.GetBucketAnalyticsConfigurationRequest 
GetBucketCorsRequest(*s3.GetBucketCorsInput) s3.GetBucketCorsRequest 
GetBucketEncryptionRequest(*s3.GetBucketEncryptionInput) s3.GetBucketEncryptionRequest 
GetBucketInventoryConfigurationRequest(*s3.GetBucketInventoryConfigurationInput) s3.GetBucketInventoryConfigurationRequest 
GetBucketLifecycleRequest(*s3.GetBucketLifecycleInput) s3.GetBucketLifecycleRequest 
GetBucketLifecycleConfigurationRequest(*s3.GetBucketLifecycleConfigurationInput) s3.GetBucketLifecycleConfigurationRequest 
GetBucketLocationRequest(*s3.GetBucketLocationInput) s3.GetBucketLocationRequest 
GetBucketLoggingRequest(*s3.GetBucketLoggingInput) s3.GetBucketLoggingRequest 
GetBucketMetricsConfigurationRequest(*s3.GetBucketMetricsConfigurationInput) s3.GetBucketMetricsConfigurationRequest 
GetBucketNotificationRequest(*s3.GetBucketNotificationConfigurationInput) s3.GetBucketNotificationRequest 
GetBucketNotificationConfigurationRequest(*s3.GetBucketNotificationConfigurationInput) s3.GetBucketNotificationConfigurationRequest 
GetBucketPolicyRequest(*s3.GetBucketPolicyInput) s3.GetBucketPolicyRequest 
GetBucketPolicyStatusRequest(*s3.GetBucketPolicyStatusInput) s3.GetBucketPolicyStatusRequest 
GetBucketReplicationRequest(*s3.GetBucketReplicationInput) s3.GetBucketReplicationRequest 
GetBucketRequestPaymentRequest(*s3.GetBucketRequestPaymentInput) s3.GetBucketRequestPaymentRequest 
GetBucketTaggingRequest(*s3.GetBucketTaggingInput) s3.GetBucketTaggingRequest 
GetBucketWebsiteRequest(*s3.GetBucketWebsiteInput) s3.GetBucketWebsiteRequest 
GetObjectAclRequest(*s3.GetObjectAclInput) s3.GetObjectAclRequest 
GetObjectLegalHoldRequest(*s3.GetObjectLegalHoldInput) s3.GetObjectLegalHoldRequest 
GetObjectLockConfigurationRequest(*s3.GetObjectLockConfigurationInput) s3.GetObjectLockConfigurationRequest 
GetObjectRetentionRequest(*s3.GetObjectRetentionInput) s3.GetObjectRetentionRequest 
GetObjectTaggingRequest(*s3.GetObjectTaggingInput) s3.GetObjectTaggingRequest 
GetObjectTorrentRequest(*s3.GetObjectTorrentInput) s3.GetObjectTorrentRequest 
GetPublicAccessBlockRequest(*s3.GetPublicAccessBlockInput) s3.GetPublicAccessBlockRequest 
HeadBucketRequest(*s3.HeadBucketInput) s3.HeadBucketRequest 
HeadObjectRequest(*s3.HeadObjectInput) s3.HeadObjectRequest 
ListBucketAnalyticsConfigurationsRequest(*s3.ListBucketAnalyticsConfigurationsInput) s3.ListBucketAnalyticsConfigurationsRequest 
ListBucketInventoryConfigurationsRequest(*s3.ListBucketInventoryConfigurationsInput) s3.ListBucketInventoryConfigurationsRequest 
ListBucketMetricsConfigurationsRequest(*s3.ListBucketMetricsConfigurationsInput) s3.ListBucketMetricsConfigurationsRequest 
ListMultipartUploadsRequest(*s3.ListMultipartUploadsInput) s3.ListMultipartUploadsRequest 
ListObjectVersionsRequest(*s3.ListObjectVersionsInput) s3.ListObjectVersionsRequest 
ListObjectsV2Request(*s3.ListObjectsV2Input) s3.ListObjectsV2Request 
ListPartsRequest(*s3.ListPartsInput) s3.ListPartsRequest 
PutBucketAccelerateConfigurationRequest(*s3.PutBucketAccelerateConfigurationInput) s3.PutBucketAccelerateConfigurationRequest 
PutBucketAclRequest(*s3.PutBucketAclInput) s3.PutBucketAclRequest 
PutBucketAnalyticsConfigurationRequest(*s3.PutBucketAnalyticsConfigurationInput) s3.PutBucketAnalyticsConfigurationRequest 
PutBucketCorsRequest(*s3.PutBucketCorsInput) s3.PutBucketCorsRequest 
PutBucketEncryptionRequest(*s3.PutBucketEncryptionInput) s3.PutBucketEncryptionRequest 
PutBucketInventoryConfigurationRequest(*s3.PutBucketInventoryConfigurationInput) s3.PutBucketInventoryConfigurationRequest 
PutBucketLifecycleRequest(*s3.PutBucketLifecycleInput) s3.PutBucketLifecycleRequest 
PutBucketLifecycleConfigurationRequest(*s3.PutBucketLifecycleConfigurationInput) s3.PutBucketLifecycleConfigurationRequest 
PutBucketLoggingRequest(*s3.PutBucketLoggingInput) s3.PutBucketLoggingRequest 
PutBucketMetricsConfigurationRequest(*s3.PutBucketMetricsConfigurationInput) s3.PutBucketMetricsConfigurationRequest 
PutBucketNotificationRequest(*s3.PutBucketNotificationInput) s3.PutBucketNotificationRequest 
PutBucketNotificationConfigurationRequest(*s3.PutBucketNotificationConfigurationInput) s3.PutBucketNotificationConfigurationRequest 
PutBucketPolicyRequest(*s3.PutBucketPolicyInput) s3.PutBucketPolicyRequest 
PutBucketReplicationRequest(*s3.PutBucketReplicationInput) s3.PutBucketReplicationRequest 
PutBucketRequestPaymentRequest(*s3.PutBucketRequestPaymentInput) s3.PutBucketRequestPaymentRequest 
PutBucketTaggingRequest(*s3.PutBucketTaggingInput) s3.PutBucketTaggingRequest 
PutBucketWebsiteRequest(*s3.PutBucketWebsiteInput) s3.PutBucketWebsiteRequest 
PutObjectAclRequest(*s3.PutObjectAclInput) s3.PutObjectAclRequest 
PutObjectLegalHoldRequest(*s3.PutObjectLegalHoldInput) s3.PutObjectLegalHoldRequest 
PutObjectLockConfigurationRequest(*s3.PutObjectLockConfigurationInput) s3.PutObjectLockConfigurationRequest 
PutObjectRetentionRequest(*s3.PutObjectRetentionInput) s3.PutObjectRetentionRequest 
PutObjectTaggingRequest(*s3.PutObjectTaggingInput) s3.PutObjectTaggingRequest 
PutPublicAccessBlockRequest(*s3.PutPublicAccessBlockInput) s3.PutPublicAccessBlockRequest 
RestoreObjectRequest(*s3.RestoreObjectInput) s3.RestoreObjectRequest 
UploadPartRequest(*s3.UploadPartInput) s3.UploadPartRequest 
UploadPartCopyRequest(*s3.UploadPartCopyInput) s3.UploadPartCopyRequest 
WaitUntilBucketExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error 
WaitUntilBucketNotExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error 
WaitUntilObjectExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error 
WaitUntilObjectNotExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error 
```

## Limitation

- Pagination is not implemented, all items are returned.
