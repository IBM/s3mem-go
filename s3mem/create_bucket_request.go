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
	"time"

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const opCreateBucket = "CreateBucket"

//CreateBucketRequest ...
func (c *Client) CreateBucketRequest(input *s3.CreateBucketInput) s3.CreateBucketRequest {
	if input == nil {
		input = &s3.CreateBucketInput{}
	}
	output := &s3.CreateBucketOutput{
		Location: input.Bucket,
	}
	op := &aws.Operation{
		Name:       opCreateBucket,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}",
	}
	req := c.NewRequest(op, input, output)
	return s3.CreateBucketRequest{Request: req, Input: input}
}

func createBucket(req *aws.Request) {
	if IsBucketExist(req.Params.(*s3.CreateBucketInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeBucketAlreadyExists, "", nil, req.Params.(*s3.CreateBucketInput).Bucket, nil, nil)
		return
	}
	tc := time.Now()
	bucket := &s3.Bucket{
		CreationDate: &tc,
		Name:         req.Params.(*s3.CreateBucketInput).Bucket,
	}
	CreateBucket(bucket)
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.CreateBucketOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
