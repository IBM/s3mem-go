package s3mem

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
)

type S3Mem struct {
}

type Buckets struct {
	Buckets map[string]*s3.Bucket
	Mux     sync.Mutex
}

//(bucket/key/version)
type Objects struct {
	Objects map[string]map[string]map[string]*Object
	Mux     sync.Mutex
}

type Object struct {
	Object  *s3.Object
	Content []byte
}

var S3MemBuckets Buckets
var S3MemObjects Objects

func init() {
	S3MemBuckets.Buckets = make(map[string]*s3.Bucket, 0)
	S3MemObjects.Objects = make(map[string]map[string]map[string]*Object, 0)
}

func Clear() {
	S3MemBuckets.Buckets = make(map[string]*s3.Bucket, 0)
	S3MemObjects.Objects = make(map[string]map[string]map[string]*Object, 0)
}

func AddBucket(b *s3.Bucket) {
	S3MemBuckets.Buckets[*b.Name] = b
}

func DeleteBucket(b *s3.Bucket) {
	delete(S3MemBuckets.Buckets, *b.Name)
}

func AddObject(bucket string, key string, body io.ReadSeeker) (*Object, error) {
	if _, ok := S3MemBuckets.Buckets[bucket]; !ok {
		return nil, errors.New(s3.ErrCodeNoSuchBucket)
	}
	if _, ok := S3MemObjects.Objects[bucket]; !ok {
		S3MemObjects.Objects[bucket] = make(map[string]map[string]*Object, 0)
	}
	if _, ok := S3MemObjects.Objects[bucket][key]; !ok {
		S3MemObjects.Objects[bucket][key] = make(map[string]*Object, 0)
	}
	tc := time.Now()
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, errors.New(s3.ErrCodeNoSuchUpload)
	}
	keyP := &key
	S3MemObjects.Objects[bucket][key]["1"] = &Object{
		Object: &s3.Object{
			Key:          keyP,
			LastModified: &tc,
			StorageClass: "memory",
		},
		Content: content,
	}
	return S3MemObjects.Objects[bucket][key]["1"], nil
}

func NewClient() (s3iface.S3API, error) {
	return &S3Mem{}, nil
}

func (c *S3Mem) AbortMultipartUploadRequest(*s3.AbortMultipartUploadInput) s3.AbortMultipartUploadRequest {
	panic("Not implemented")
	return s3.AbortMultipartUploadRequest{}
}

func (c *S3Mem) CompleteMultipartUploadRequest(*s3.CompleteMultipartUploadInput) s3.CompleteMultipartUploadRequest {
	panic("Not implemented")
	return s3.CompleteMultipartUploadRequest{}
}

func (c *S3Mem) CopyObjectRequest(*s3.CopyObjectInput) s3.CopyObjectRequest {
	panic("Not implemented")
	return s3.CopyObjectRequest{}
}

func (c *S3Mem) CreateBucketRequest(input *s3.CreateBucketInput) s3.CreateBucketRequest {
	if input == nil {
		input = &s3.CreateBucketInput{}
	}
	output := &s3.CreateBucketOutput{
		Location: input.Bucket,
	}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	tc := time.Now()
	S3MemBuckets.Buckets[*input.Bucket] = &s3.Bucket{
		CreationDate: &tc,
		Name:         input.Bucket,
	}
	return s3.CreateBucketRequest{Request: req, Input: input, Copy: c.CreateBucketRequest}
}

func (c *S3Mem) CreateMultipartUploadRequest(*s3.CreateMultipartUploadInput) s3.CreateMultipartUploadRequest {
	panic("Not implemented")
	return s3.CreateMultipartUploadRequest{}
}

func (c *S3Mem) DeleteBucketRequest(input *s3.DeleteBucketInput) s3.DeleteBucketRequest {
	if input == nil {
		input = &s3.DeleteBucketInput{}
	}
	output := &s3.DeleteBucketOutput{}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	if _, ok := S3MemBuckets.Buckets[*input.Bucket]; !ok {
		req.Error = errors.New(s3.ErrCodeNoSuchBucket)
	}
	delete(S3MemBuckets.Buckets, *input.Bucket)
	return s3.DeleteBucketRequest{Request: req, Input: input, Copy: c.DeleteBucketRequest}
}

func (c *S3Mem) DeleteBucketAnalyticsConfigurationRequest(*s3.DeleteBucketAnalyticsConfigurationInput) s3.DeleteBucketAnalyticsConfigurationRequest {
	panic("Not implemented")
	return s3.DeleteBucketAnalyticsConfigurationRequest{}
}

func (c *S3Mem) DeleteBucketCorsRequest(*s3.DeleteBucketCorsInput) s3.DeleteBucketCorsRequest {
	panic("Not implemented")
	return s3.DeleteBucketCorsRequest{}
}

func (c *S3Mem) DeleteBucketEncryptionRequest(*s3.DeleteBucketEncryptionInput) s3.DeleteBucketEncryptionRequest {
	panic("Not implemented")
	return s3.DeleteBucketEncryptionRequest{}
}

func (c *S3Mem) DeleteBucketInventoryConfigurationRequest(*s3.DeleteBucketInventoryConfigurationInput) s3.DeleteBucketInventoryConfigurationRequest {
	panic("Not implemented")
	return s3.DeleteBucketInventoryConfigurationRequest{}
}

func (c *S3Mem) DeleteBucketLifecycleRequest(*s3.DeleteBucketLifecycleInput) s3.DeleteBucketLifecycleRequest {
	panic("Not implemented")
	return s3.DeleteBucketLifecycleRequest{}
}

func (c *S3Mem) DeleteBucketMetricsConfigurationRequest(*s3.DeleteBucketMetricsConfigurationInput) s3.DeleteBucketMetricsConfigurationRequest {
	panic("Not implemented")
	return s3.DeleteBucketMetricsConfigurationRequest{}
}

func (c *S3Mem) DeleteBucketPolicyRequest(*s3.DeleteBucketPolicyInput) s3.DeleteBucketPolicyRequest {
	panic("Not implemented")
	return s3.DeleteBucketPolicyRequest{}
}

func (c *S3Mem) DeleteBucketReplicationRequest(*s3.DeleteBucketReplicationInput) s3.DeleteBucketReplicationRequest {
	panic("Not implemented")
	return s3.DeleteBucketReplicationRequest{}
}

func (c *S3Mem) DeleteBucketTaggingRequest(*s3.DeleteBucketTaggingInput) s3.DeleteBucketTaggingRequest {
	panic("Not implemented")
	return s3.DeleteBucketTaggingRequest{}
}

func (c *S3Mem) DeleteBucketWebsiteRequest(*s3.DeleteBucketWebsiteInput) s3.DeleteBucketWebsiteRequest {
	panic("Not implemented")
	return s3.DeleteBucketWebsiteRequest{}
}

func (c *S3Mem) DeleteObjectRequest(input *s3.DeleteObjectInput) s3.DeleteObjectRequest {
	if input == nil {
		input = &s3.DeleteObjectInput{}
	}
	output := &s3.DeleteObjectOutput{}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	if _, ok := S3MemObjects.Objects[*input.Bucket]; !ok {
		req.Error = errors.New(s3.ErrCodeNoSuchBucket)
		return s3.DeleteObjectRequest{Request: req, Input: input, Copy: c.DeleteObjectRequest}
	}
	if _, ok := S3MemObjects.Objects[*input.Bucket][*input.Key]; !ok {
		req.Error = errors.New(s3.ErrCodeNoSuchKey)
		return s3.DeleteObjectRequest{Request: req, Input: input, Copy: c.DeleteObjectRequest}
	}
	delete(S3MemObjects.Objects[*input.Bucket][*input.Key], "1")
	delete(S3MemObjects.Objects[*input.Bucket], *input.Key)
	return s3.DeleteObjectRequest{Request: req, Input: input, Copy: c.DeleteObjectRequest}
}

func (c *S3Mem) DeleteObjectTaggingRequest(*s3.DeleteObjectTaggingInput) s3.DeleteObjectTaggingRequest {
	panic("Not implemented")
	return s3.DeleteObjectTaggingRequest{}
}

func (c *S3Mem) DeleteObjectsRequest(*s3.DeleteObjectsInput) s3.DeleteObjectsRequest {
	panic("Not implemented")
	return s3.DeleteObjectsRequest{}
}

func (c *S3Mem) DeletePublicAccessBlockRequest(*s3.DeletePublicAccessBlockInput) s3.DeletePublicAccessBlockRequest {
	panic("Not implemented")
	return s3.DeletePublicAccessBlockRequest{}
}

func (c *S3Mem) GetBucketAccelerateConfigurationRequest(*s3.GetBucketAccelerateConfigurationInput) s3.GetBucketAccelerateConfigurationRequest {
	panic("Not implemented")
	return s3.GetBucketAccelerateConfigurationRequest{}
}

func (c *S3Mem) GetBucketAclRequest(*s3.GetBucketAclInput) s3.GetBucketAclRequest {
	panic("Not implemented")
	return s3.GetBucketAclRequest{}
}

func (c *S3Mem) GetBucketAnalyticsConfigurationRequest(*s3.GetBucketAnalyticsConfigurationInput) s3.GetBucketAnalyticsConfigurationRequest {
	panic("Not implemented")
	return s3.GetBucketAnalyticsConfigurationRequest{}
}

func (c *S3Mem) GetBucketCorsRequest(*s3.GetBucketCorsInput) s3.GetBucketCorsRequest {
	panic("Not implemented")
	return s3.GetBucketCorsRequest{}
}

func (c *S3Mem) GetBucketEncryptionRequest(*s3.GetBucketEncryptionInput) s3.GetBucketEncryptionRequest {
	panic("Not implemented")
	return s3.GetBucketEncryptionRequest{}
}

func (c *S3Mem) GetBucketInventoryConfigurationRequest(*s3.GetBucketInventoryConfigurationInput) s3.GetBucketInventoryConfigurationRequest {
	panic("Not implemented")
	return s3.GetBucketInventoryConfigurationRequest{}
}

func (c *S3Mem) GetBucketLifecycleRequest(*s3.GetBucketLifecycleInput) s3.GetBucketLifecycleRequest {
	panic("Not implemented")
	return s3.GetBucketLifecycleRequest{}
}

func (c *S3Mem) GetBucketLifecycleConfigurationRequest(*s3.GetBucketLifecycleConfigurationInput) s3.GetBucketLifecycleConfigurationRequest {
	panic("Not implemented")
	return s3.GetBucketLifecycleConfigurationRequest{}
}

func (c *S3Mem) GetBucketLocationRequest(*s3.GetBucketLocationInput) s3.GetBucketLocationRequest {
	panic("Not implemented")
	return s3.GetBucketLocationRequest{}
}

func (c *S3Mem) GetBucketLoggingRequest(*s3.GetBucketLoggingInput) s3.GetBucketLoggingRequest {
	panic("Not implemented")
	return s3.GetBucketLoggingRequest{}
}

func (c *S3Mem) GetBucketMetricsConfigurationRequest(*s3.GetBucketMetricsConfigurationInput) s3.GetBucketMetricsConfigurationRequest {
	panic("Not implemented")
	return s3.GetBucketMetricsConfigurationRequest{}
}

func (c *S3Mem) GetBucketNotificationRequest(*s3.GetBucketNotificationConfigurationInput) s3.GetBucketNotificationRequest {
	panic("Not implemented")
	return s3.GetBucketNotificationRequest{}
}

func (c *S3Mem) GetBucketNotificationConfigurationRequest(*s3.GetBucketNotificationConfigurationInput) s3.GetBucketNotificationConfigurationRequest {
	panic("Not implemented")
	return s3.GetBucketNotificationConfigurationRequest{}
}

func (c *S3Mem) GetBucketPolicyRequest(*s3.GetBucketPolicyInput) s3.GetBucketPolicyRequest {
	panic("Not implemented")
	return s3.GetBucketPolicyRequest{}
}

func (c *S3Mem) GetBucketPolicyStatusRequest(*s3.GetBucketPolicyStatusInput) s3.GetBucketPolicyStatusRequest {
	panic("Not implemented")
	return s3.GetBucketPolicyStatusRequest{}
}

func (c *S3Mem) GetBucketReplicationRequest(*s3.GetBucketReplicationInput) s3.GetBucketReplicationRequest {
	panic("Not implemented")
	return s3.GetBucketReplicationRequest{}
}

func (c *S3Mem) GetBucketRequestPaymentRequest(*s3.GetBucketRequestPaymentInput) s3.GetBucketRequestPaymentRequest {
	panic("Not implemented")
	return s3.GetBucketRequestPaymentRequest{}
}

func (c *S3Mem) GetBucketTaggingRequest(*s3.GetBucketTaggingInput) s3.GetBucketTaggingRequest {
	panic("Not implemented")
	return s3.GetBucketTaggingRequest{}
}

func (c *S3Mem) GetBucketVersioningRequest(*s3.GetBucketVersioningInput) s3.GetBucketVersioningRequest {
	panic("Not implemented")
	return s3.GetBucketVersioningRequest{}
}

func (c *S3Mem) GetBucketWebsiteRequest(*s3.GetBucketWebsiteInput) s3.GetBucketWebsiteRequest {
	panic("Not implemented")
	return s3.GetBucketWebsiteRequest{}
}

func (c *S3Mem) GetObjectRequest(input *s3.GetObjectInput) s3.GetObjectRequest {
	if input == nil {
		input = &s3.GetObjectInput{}
	}
	output := &s3.GetObjectOutput{}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	if _, ok := S3MemObjects.Objects[*input.Bucket]; !ok {
		req.Error = errors.New(s3.ErrCodeNoSuchBucket)
		return s3.GetObjectRequest{Request: req, Input: input, Copy: c.GetObjectRequest}
	}
	if _, ok := S3MemObjects.Objects[*input.Bucket][*input.Key]; !ok {
		req.Error = errors.New(s3.ErrCodeNoSuchKey)
		return s3.GetObjectRequest{Request: req, Input: input, Copy: c.GetObjectRequest}
	}
	output.Body = ioutil.NopCloser(bytes.NewReader(S3MemObjects.Objects[*input.Bucket][*input.Key]["1"].Content))
	return s3.GetObjectRequest{Request: req, Input: input, Copy: c.GetObjectRequest}
}

func (c *S3Mem) GetObjectAclRequest(*s3.GetObjectAclInput) s3.GetObjectAclRequest {
	panic("Not implemented")
	return s3.GetObjectAclRequest{}
}

func (c *S3Mem) GetObjectLegalHoldRequest(*s3.GetObjectLegalHoldInput) s3.GetObjectLegalHoldRequest {
	panic("Not implemented")
	return s3.GetObjectLegalHoldRequest{}
}

func (c *S3Mem) GetObjectLockConfigurationRequest(*s3.GetObjectLockConfigurationInput) s3.GetObjectLockConfigurationRequest {
	panic("Not implemented")
	return s3.GetObjectLockConfigurationRequest{}
}

func (c *S3Mem) GetObjectRetentionRequest(*s3.GetObjectRetentionInput) s3.GetObjectRetentionRequest {
	panic("Not implemented")
	return s3.GetObjectRetentionRequest{}
}

func (c *S3Mem) GetObjectTaggingRequest(*s3.GetObjectTaggingInput) s3.GetObjectTaggingRequest {
	panic("Not implemented")
	return s3.GetObjectTaggingRequest{}
}

func (c *S3Mem) GetObjectTorrentRequest(*s3.GetObjectTorrentInput) s3.GetObjectTorrentRequest {
	panic("Not implemented")
	return s3.GetObjectTorrentRequest{}
}

func (c *S3Mem) GetPublicAccessBlockRequest(*s3.GetPublicAccessBlockInput) s3.GetPublicAccessBlockRequest {
	panic("Not implemented")
	return s3.GetPublicAccessBlockRequest{}
}

func (c *S3Mem) HeadBucketRequest(*s3.HeadBucketInput) s3.HeadBucketRequest {
	panic("Not implemented")
	return s3.HeadBucketRequest{}
}

func (c *S3Mem) HeadObjectRequest(*s3.HeadObjectInput) s3.HeadObjectRequest {
	panic("Not implemented")
	return s3.HeadObjectRequest{}
}

func (c *S3Mem) ListBucketAnalyticsConfigurationsRequest(*s3.ListBucketAnalyticsConfigurationsInput) s3.ListBucketAnalyticsConfigurationsRequest {
	panic("Not implemented")
	return s3.ListBucketAnalyticsConfigurationsRequest{}
}

func (c *S3Mem) ListBucketInventoryConfigurationsRequest(*s3.ListBucketInventoryConfigurationsInput) s3.ListBucketInventoryConfigurationsRequest {
	panic("Not implemented")
	return s3.ListBucketInventoryConfigurationsRequest{}
}

func (c *S3Mem) ListBucketMetricsConfigurationsRequest(*s3.ListBucketMetricsConfigurationsInput) s3.ListBucketMetricsConfigurationsRequest {
	panic("Not implemented")
	return s3.ListBucketMetricsConfigurationsRequest{}
}

func (c *S3Mem) ListBucketsRequest(input *s3.ListBucketsInput) s3.ListBucketsRequest {
	if input == nil {
		input = &s3.ListBucketsInput{}
	}
	v := make([]s3.Bucket, 0)
	for _, value := range S3MemBuckets.Buckets {
		v = append(v, *value)
	}
	output := &s3.ListBucketsOutput{
		Buckets: v,
	}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	return s3.ListBucketsRequest{Request: req, Input: input, Copy: c.ListBucketsRequest}
}

func (c *S3Mem) ListMultipartUploadsRequest(*s3.ListMultipartUploadsInput) s3.ListMultipartUploadsRequest {
	panic("Not implemented")
	return s3.ListMultipartUploadsRequest{}
}

func (c *S3Mem) ListObjectVersionsRequest(*s3.ListObjectVersionsInput) s3.ListObjectVersionsRequest {
	panic("Not implemented")
	return s3.ListObjectVersionsRequest{}
}

func (c *S3Mem) ListObjectsRequest(input *s3.ListObjectsInput) s3.ListObjectsRequest {
	if input == nil {
		input = &s3.ListObjectsInput{}
	}
	output := &s3.ListObjectsOutput{}
	operation := &aws.Operation{}
	req := &aws.Request{
		Data:        output,
		Operation:   operation,
		HTTPRequest: &http.Request{},
	}
	if _, ok := S3MemBuckets.Buckets[*input.Bucket]; !ok {
		req.Error = errors.New(s3.ErrCodeNoSuchBucket)
	}
	v := make([]s3.Object, 0)
	for _, obj := range S3MemObjects.Objects[*input.Bucket] {
		if strings.HasPrefix(*obj["1"].Object.Key, *input.Prefix) {
			v = append(v, *obj["1"].Object)
		}
	}
	output.Contents = v
	return s3.ListObjectsRequest{Request: req, Input: input, Copy: c.ListObjectsRequest}
}

func (c *S3Mem) ListObjectsV2Request(*s3.ListObjectsV2Input) s3.ListObjectsV2Request {
	panic("Not implemented")
	return s3.ListObjectsV2Request{}
}

func (c *S3Mem) ListPartsRequest(*s3.ListPartsInput) s3.ListPartsRequest {
	panic("Not implemented")
	return s3.ListPartsRequest{}
}

func (c *S3Mem) PutBucketAccelerateConfigurationRequest(*s3.PutBucketAccelerateConfigurationInput) s3.PutBucketAccelerateConfigurationRequest {
	panic("Not implemented")
	return s3.PutBucketAccelerateConfigurationRequest{}
}

func (c *S3Mem) PutBucketAclRequest(*s3.PutBucketAclInput) s3.PutBucketAclRequest {
	panic("Not implemented")
	return s3.PutBucketAclRequest{}
}

func (c *S3Mem) PutBucketAnalyticsConfigurationRequest(*s3.PutBucketAnalyticsConfigurationInput) s3.PutBucketAnalyticsConfigurationRequest {
	panic("Not implemented")
	return s3.PutBucketAnalyticsConfigurationRequest{}
}

func (c *S3Mem) PutBucketCorsRequest(*s3.PutBucketCorsInput) s3.PutBucketCorsRequest {
	panic("Not implemented")
	return s3.PutBucketCorsRequest{}
}

func (c *S3Mem) PutBucketEncryptionRequest(*s3.PutBucketEncryptionInput) s3.PutBucketEncryptionRequest {
	panic("Not implemented")
	return s3.PutBucketEncryptionRequest{}
}

func (c *S3Mem) PutBucketInventoryConfigurationRequest(*s3.PutBucketInventoryConfigurationInput) s3.PutBucketInventoryConfigurationRequest {
	panic("Not implemented")
	return s3.PutBucketInventoryConfigurationRequest{}
}

func (c *S3Mem) PutBucketLifecycleRequest(*s3.PutBucketLifecycleInput) s3.PutBucketLifecycleRequest {
	panic("Not implemented")
	return s3.PutBucketLifecycleRequest{}
}

func (c *S3Mem) PutBucketLifecycleConfigurationRequest(*s3.PutBucketLifecycleConfigurationInput) s3.PutBucketLifecycleConfigurationRequest {
	panic("Not implemented")
	return s3.PutBucketLifecycleConfigurationRequest{}
}

func (c *S3Mem) PutBucketLoggingRequest(*s3.PutBucketLoggingInput) s3.PutBucketLoggingRequest {
	panic("Not implemented")
	return s3.PutBucketLoggingRequest{}
}

func (c *S3Mem) PutBucketMetricsConfigurationRequest(*s3.PutBucketMetricsConfigurationInput) s3.PutBucketMetricsConfigurationRequest {
	panic("Not implemented")
	return s3.PutBucketMetricsConfigurationRequest{}
}

func (c *S3Mem) PutBucketNotificationRequest(*s3.PutBucketNotificationInput) s3.PutBucketNotificationRequest {
	panic("Not implemented")
	return s3.PutBucketNotificationRequest{}
}

func (c *S3Mem) PutBucketNotificationConfigurationRequest(*s3.PutBucketNotificationConfigurationInput) s3.PutBucketNotificationConfigurationRequest {
	panic("Not implemented")
	return s3.PutBucketNotificationConfigurationRequest{}
}

func (c *S3Mem) PutBucketPolicyRequest(*s3.PutBucketPolicyInput) s3.PutBucketPolicyRequest {
	panic("Not implemented")
	return s3.PutBucketPolicyRequest{}
}

func (c *S3Mem) PutBucketReplicationRequest(*s3.PutBucketReplicationInput) s3.PutBucketReplicationRequest {
	panic("Not implemented")
	return s3.PutBucketReplicationRequest{}
}

func (c *S3Mem) PutBucketRequestPaymentRequest(*s3.PutBucketRequestPaymentInput) s3.PutBucketRequestPaymentRequest {
	panic("Not implemented")
	return s3.PutBucketRequestPaymentRequest{}
}

func (c *S3Mem) PutBucketTaggingRequest(*s3.PutBucketTaggingInput) s3.PutBucketTaggingRequest {
	panic("Not implemented")
	return s3.PutBucketTaggingRequest{}
}

func (c *S3Mem) PutBucketVersioningRequest(*s3.PutBucketVersioningInput) s3.PutBucketVersioningRequest {
	panic("Not implemented")
	return s3.PutBucketVersioningRequest{}
}

func (c *S3Mem) PutBucketWebsiteRequest(*s3.PutBucketWebsiteInput) s3.PutBucketWebsiteRequest {
	panic("Not implemented")
	return s3.PutBucketWebsiteRequest{}
}

func (c *S3Mem) PutObjectRequest(input *s3.PutObjectInput) s3.PutObjectRequest {
	if input == nil {
		input = &s3.PutObjectInput{}
	}
	output := &s3.PutObjectOutput{}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	_, err := AddObject(*input.Bucket, *input.Key, input.Body)
	if err != nil {
		req.Error = errors.New(s3.ErrCodeNoSuchUpload)
		return s3.PutObjectRequest{Request: req, Input: input, Copy: c.PutObjectRequest}
	}
	return s3.PutObjectRequest{Request: req, Input: input, Copy: c.PutObjectRequest}
}

func (c *S3Mem) PutObjectAclRequest(*s3.PutObjectAclInput) s3.PutObjectAclRequest {
	panic("Not implemented")
	return s3.PutObjectAclRequest{}
}

func (c *S3Mem) PutObjectLegalHoldRequest(*s3.PutObjectLegalHoldInput) s3.PutObjectLegalHoldRequest {
	panic("Not implemented")
	return s3.PutObjectLegalHoldRequest{}
}

func (c *S3Mem) PutObjectLockConfigurationRequest(*s3.PutObjectLockConfigurationInput) s3.PutObjectLockConfigurationRequest {
	panic("Not implemented")
	return s3.PutObjectLockConfigurationRequest{}
}

func (c *S3Mem) PutObjectRetentionRequest(*s3.PutObjectRetentionInput) s3.PutObjectRetentionRequest {
	panic("Not implemented")
	return s3.PutObjectRetentionRequest{}
}

func (c *S3Mem) PutObjectTaggingRequest(*s3.PutObjectTaggingInput) s3.PutObjectTaggingRequest {
	panic("Not implemented")
	return s3.PutObjectTaggingRequest{}
}

func (c *S3Mem) PutPublicAccessBlockRequest(*s3.PutPublicAccessBlockInput) s3.PutPublicAccessBlockRequest {
	panic("Not implemented")
	return s3.PutPublicAccessBlockRequest{}
}

func (c *S3Mem) RestoreObjectRequest(*s3.RestoreObjectInput) s3.RestoreObjectRequest {
	panic("Not implemented")
	return s3.RestoreObjectRequest{}
}

func (c *S3Mem) UploadPartRequest(*s3.UploadPartInput) s3.UploadPartRequest {
	panic("Not implemented")
	return s3.UploadPartRequest{}
}

func (c *S3Mem) UploadPartCopyRequest(*s3.UploadPartCopyInput) s3.UploadPartCopyRequest {
	panic("Not implemented")
	return s3.UploadPartCopyRequest{}
}

func (c *S3Mem) WaitUntilBucketExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error {
	panic("Not implemented")
	return nil
}

func (c *S3Mem) WaitUntilBucketNotExists(context.Context, *s3.HeadBucketInput, ...aws.WaiterOption) error {
	panic("Not implemented")
	return nil
}

func (c *S3Mem) WaitUntilObjectExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error {
	panic("Not implemented")
	return nil
}

func (c *S3Mem) WaitUntilObjectNotExists(context.Context, *s3.HeadObjectInput, ...aws.WaiterOption) error {
	panic("Not implemented")
	return nil
}
