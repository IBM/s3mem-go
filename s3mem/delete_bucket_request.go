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
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

//DeleteBucketRequest ...
func (c *S3Mem) DeleteBucketRequest(input *s3.DeleteBucketInput) s3.DeleteBucketRequest {
	if input == nil {
		input = &s3.DeleteBucketInput{}
	}
	output := &s3.DeleteBucketOutput{}
	req := &aws.Request{
		Params:      input,
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	bucketExists := aws.NamedHandler{Name: "S3MemBucketExists", Fn: deleteBucketBucketExists}
	req.Handlers.Send.PushBackNamed(bucketExists)
	bucketIsEmpty := aws.NamedHandler{Name: "S3MemBucketIsEmpty", Fn: deleteBucketBucketIsEmpty}
	req.Handlers.Send.PushBackNamed(bucketIsEmpty)
	deleteBucket := aws.NamedHandler{Name: "S3MemDeleteBucket", Fn: deleteBucket}
	req.Handlers.Send.PushBackNamed(deleteBucket)
	return s3.DeleteBucketRequest{Request: req, Input: input}
}

func deleteBucketBucketExists(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsBucketExist(req.Params.(*s3.DeleteBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteBucketInput).Bucket, nil, nil)
	}
}

func deleteBucketBucketIsEmpty(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsBucketEmpty(req.Params.(*s3.DeleteBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3memerr.ErrCodeBucketNotEmpty, "", nil, req.Params.(*s3.DeleteBucketInput).Bucket, nil, nil)
	}
}

func deleteBucket(req *aws.Request) {
	if req.Error != nil {
		return
	}
	DeleteBucket(req.Params.(*s3.DeleteBucketInput).Bucket)
}
