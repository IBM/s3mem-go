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
