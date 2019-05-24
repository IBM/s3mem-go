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
	"errors"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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
