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
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

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
