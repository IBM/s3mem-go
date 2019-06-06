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
	"time"

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *S3Mem) CreateBucketRequest(input *s3.CreateBucketInput) s3.CreateBucketRequest {
	if input == nil {
		input = &s3.CreateBucketInput{}
	}
	output := &s3.CreateBucketOutput{
		Location: input.Bucket,
	}
	req := &aws.Request{
		Params:      input,
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	createBucketNameValidate := aws.NamedHandler{Name: "S3MemCreateBucketNameValidate", Fn: createBucketBucketExists}
	req.Handlers.Send.PushBackNamed(createBucketNameValidate)
	addBucketNameSend := aws.NamedHandler{Name: "S3MemCreateBucketNameSend", Fn: createBucket}
	req.Handlers.Send.PushBackNamed(addBucketNameSend)
	return s3.CreateBucketRequest{Request: req, Input: input, Copy: c.CreateBucketRequest}
}

func createBucketBucketExists(req *aws.Request) {
	if IsBucketExist(req.Params.(*s3.CreateBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeBucketAlreadyExists, "", nil, req.Params.(*s3.CreateBucketInput).Bucket, nil, nil)
	}
}

func createBucket(req *aws.Request) {
	tc := time.Now()
	bucket := &s3.Bucket{
		CreationDate: &tc,
		Name:         req.Params.(*s3.CreateBucketInput).Bucket,
	}
	CreateBucket(bucket)
}
