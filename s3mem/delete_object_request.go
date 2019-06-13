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

const opDeleteObject = "DeleteObject"

//DeleteObjectRequest ...
func (c *Client) DeleteObjectRequest(input *s3.DeleteObjectInput) s3.DeleteObjectRequest {
	if input == nil {
		input = &s3.DeleteObjectInput{}
	}
	output := &s3.DeleteObjectOutput{}
	op := &aws.Operation{
		Name:       opDeleteObject,
		HTTPMethod: "DELETE",
		HTTPPath:   "/{Bucket}/{Key+}",
	}
	req := c.NewRequest(op, input, output)
	return s3.DeleteObjectRequest{Request: req, Input: input}
}

func deleteObject(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.DeleteObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteObjectInput).Bucket, nil, nil)
		return
	}
	if !IsObjectExist(req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key, req.Params.(*s3.DeleteObjectInput).VersionId)
		return
	}
	deleteMarker, versionID, err := DeleteObject(req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key, req.Params.(*s3.DeleteObjectInput).VersionId)
	if err != nil {
		req.Error = err
		return
	}
	req.Data.(*s3.DeleteObjectOutput).DeleteMarker = deleteMarker
	req.Data.(*s3.DeleteObjectOutput).VersionId = versionID
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.DeleteObjectOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
