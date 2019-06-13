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
	"bytes"
	"encoding/json"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

const opDeleteBucket = "DeleteBucket"

//DeleteBucketRequest ...
func (c *Client) DeleteBucketRequest(input *s3.DeleteBucketInput) s3.DeleteBucketRequest {
	if input == nil {
		input = &s3.DeleteBucketInput{}
	}
	output := &s3.DeleteBucketOutput{}
	op := &aws.Operation{
		Name:       opDeleteBucket,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}",
	}
	req := c.NewRequest(op, input, output)
	return s3.DeleteBucketRequest{Request: req, Input: input}
}

func deleteBucket(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.DeleteBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteBucketInput).Bucket, nil, nil)
		return
	}
	if !IsBucketEmpty(req.Params.(*s3.DeleteBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3memerr.ErrCodeBucketNotEmpty, "", nil, req.Params.(*s3.DeleteBucketInput).Bucket, nil, nil)
		return
	}
	DeleteBucket(req.Params.(*s3.DeleteBucketInput).Bucket)
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.DeleteBucketOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
