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

const opPutObject = "PutObject"

//PutObjectRequest ...
func (c *Client) PutObjectRequest(input *s3.PutObjectInput) s3.PutObjectRequest {
	if input == nil {
		input = &s3.PutObjectInput{}
	}
	output := &s3.PutObjectOutput{}
	op := &aws.Operation{
		Name:       opPutObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}
	req := c.NewRequest(op, input, output)
	return s3.PutObjectRequest{Request: req, Input: input}
}

func putObject(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.PutObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.PutObjectInput).Bucket, nil, nil)
		return
	}
	_, _, err := PutObject(req.Params.(*s3.PutObjectInput).Bucket, req.Params.(*s3.PutObjectInput).Key, req.Params.(*s3.PutObjectInput).Body)
	if err != nil {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchUpload, "", nil, req.Params.(*s3.PutObjectInput).Bucket, req.Params.(*s3.PutObjectInput).Key, nil)
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.PutObjectOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
