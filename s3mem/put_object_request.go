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

func (c *S3Mem) PutObjectRequest(input *s3.PutObjectInput) s3.PutObjectRequest {
	if input == nil {
		input = &s3.PutObjectInput{}
	}
	output := &s3.PutObjectOutput{}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	_, err := AddObject(input.Bucket, input.Key, input.Body)
	if err != nil {
		req.Error = errors.New(s3.ErrCodeNoSuchUpload)
		return s3.PutObjectRequest{Request: req, Input: input, Copy: c.PutObjectRequest}
	}
	return s3.PutObjectRequest{Request: req, Input: input, Copy: c.PutObjectRequest}
}
