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

func (c *S3Mem) DeleteObjectsRequest(input *s3.DeleteObjectsInput) s3.DeleteObjectsRequest {
	if input == nil {
		input = &s3.DeleteObjectsInput{}
	}
	output := &s3.DeleteObjectsOutput{}
	req := &aws.Request{
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	if _, ok := S3MemObjects.Objects[*input.Bucket]; !ok {
		req.Error = errors.New(s3.ErrCodeNoSuchBucket)
		return s3.DeleteObjectsRequest{Request: req, Input: input, Copy: c.DeleteObjectsRequest}
	}
	for _, obj := range input.Delete.Objects {
		delete(S3MemObjects.Objects[*input.Bucket][*obj.Key], "1")
		delete(S3MemObjects.Objects[*input.Bucket], *obj.Key)
	}
	return s3.DeleteObjectsRequest{Request: req, Input: input, Copy: c.DeleteObjectsRequest}
}
