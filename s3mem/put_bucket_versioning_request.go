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

func (c *S3Mem) PutBucketVersioningRequest(input *s3.PutBucketVersioningInput) s3.PutBucketVersioningRequest {
	if input == nil {
		input = &s3.PutBucketVersioningInput{}
	}
	output := &s3.PutBucketVersioningOutput{}
	req := &aws.Request{
		Params:      input,
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	bucketExists := aws.NamedHandler{Name: "S3MemBucketExists", Fn: putBucketVersioningBucketExists}
	req.Handlers.Send.PushBackNamed(bucketExists)
	putBucketVersioning := aws.NamedHandler{Name: "S3MemPutBucketVersioning", Fn: putBucketVersioning}
	req.Handlers.Send.PushBackNamed(putBucketVersioning)
	return s3.PutBucketVersioningRequest{Request: req, Input: input, Copy: c.PutBucketVersioningRequest}
}

func putBucketVersioningBucketExists(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsBucketExist(req.Params.(*s3.PutBucketVersioningInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.PutBucketVersioningInput).Bucket, nil, nil)
	}
}

func putBucketVersioning(req *aws.Request) {
	if req.Error != nil {
		return
	}
	err := PutBucketVersioning(req.Params.(*s3.PutBucketVersioningInput).Bucket, req.Params.(*s3.PutBucketVersioningInput).MFA, req.Params.(*s3.PutBucketVersioningInput).VersioningConfiguration)
	if err != nil {
		req.Error = s3memerr.NewError(s3memerr.ErrCodeIllegalVersioningConfigurationException, "", nil, req.Params.(*s3.PutBucketVersioningInput).Bucket, nil, nil)
	}
}
