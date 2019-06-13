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

const opPutBucketVersioning = "PutBucketVersioning"

//PutBucketVersioningRequest ...
func (c *Client) PutBucketVersioningRequest(input *s3.PutBucketVersioningInput) s3.PutBucketVersioningRequest {
	if input == nil {
		input = &s3.PutBucketVersioningInput{}
	}
	output := &s3.PutBucketVersioningOutput{}
	op := &aws.Operation{
		Name:       opPutBucketVersioning,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}?versioning",
	}
	req := c.NewRequest(op, input, output)
	return s3.PutBucketVersioningRequest{Request: req, Input: input}
}

func putBucketVersioningBucketExists(req *aws.Request) {
}

func putBucketVersioning(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.PutBucketVersioningInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.PutBucketVersioningInput).Bucket, nil, nil)
		return
	}
	err := PutBucketVersioning(req.Params.(*s3.PutBucketVersioningInput).Bucket, req.Params.(*s3.PutBucketVersioningInput).MFA, req.Params.(*s3.PutBucketVersioningInput).VersioningConfiguration)
	if err != nil {
		req.Error = s3memerr.NewError(s3memerr.ErrCodeIllegalVersioningConfigurationException, "", nil, req.Params.(*s3.PutBucketVersioningInput).Bucket, nil, nil)
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.PutBucketVersioningOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
