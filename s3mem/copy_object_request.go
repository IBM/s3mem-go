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
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

const opCopyObject = "CopyObject"

//CopyObjectRequest ...
func (c *Client) CopyObjectRequest(input *s3.CopyObjectInput) s3.CopyObjectRequest {
	if input == nil {
		input = &s3.CopyObjectInput{}
	}
	output := &s3.CopyObjectOutput{}
	operation := &aws.Operation{
		Name:       opCopyObject,
		HTTPMethod: "PUT",
		HTTPPath:   "/{Bucket}/{Key+}",
	}
	req := c.NewRequest(operation, input, output)
	return s3.CopyObjectRequest{Request: req, Input: input, Copy: nil}
}

func copyObject(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.CopyObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.CopyObjectInput).Bucket, nil, nil)
		return
	}
	bucket, key, err := ParseObjectURL(req.Params.(*s3.CopyObjectInput).CopySource)
	obj, versionId, err := GetObject(bucket, key, nil)
	if err != nil {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, nil)
	}
	objDest, destVersionId, err := PutObject(req.Params.(*s3.CopyObjectInput).Bucket, key, strings.NewReader(string(obj.Content)))
	if err != nil {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchUpload, "", nil, req.Params.(*s3.CopyObjectInput).Bucket, key, nil)
	}
	req.Data.(*s3.CopyObjectOutput).CopyObjectResult = &s3.CopyObjectResult{
		ETag:         objDest.Object.ETag,
		LastModified: objDest.Object.LastModified,
	}
	req.Data.(*s3.CopyObjectOutput).CopySourceVersionId = versionId
	req.Data.(*s3.CopyObjectOutput).VersionId = destVersionId
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.CopyObjectOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))

}