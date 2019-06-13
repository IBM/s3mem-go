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

const opGetObject = "GetObject"

//GetObjectRequest ...
func (c *Client) GetObjectRequest(input *s3.GetObjectInput) s3.GetObjectRequest {
	if input == nil {
		input = &s3.GetObjectInput{}
	}
	output := &s3.GetObjectOutput{}
	operation := &aws.Operation{
		Name:       opGetObject,
		HTTPMethod: "GET",
		HTTPPath:   "/{Bucket}/{Key+}",
	}
	req := c.NewRequest(operation, input, output)
	return s3.GetObjectRequest{Request: req, Input: input}
}

func getObject(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.GetObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.GetObjectInput).Bucket, nil, nil)
	}
	obj, versionId, err := GetObject(req.Params.(*s3.GetObjectInput).Bucket, req.Params.(*s3.GetObjectInput).Key, req.Params.(*s3.GetObjectInput).VersionId)
	if err != nil {
		req.Error = err
		return
	}
	if obj == nil {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, req.Params.(*s3.GetObjectInput).Bucket, req.Params.(*s3.GetObjectInput).Key, req.Params.(*s3.GetObjectInput).VersionId)
		return
	}
	req.Data.(*s3.GetObjectOutput).Body = ioutil.NopCloser(bytes.NewReader(obj.Content))
	req.Data.(*s3.GetObjectOutput).VersionId = versionId
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.GetObjectOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
