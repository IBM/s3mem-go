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
