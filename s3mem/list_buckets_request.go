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
)

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
