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
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

func (c *S3Mem) GetObjectRequest(input *s3.GetObjectInput) s3.GetObjectRequest {
	if input == nil {
		input = &s3.GetObjectInput{}
	}
	output := &s3.GetObjectOutput{}
	req := &aws.Request{
		Params:      input,
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	bucketExists := aws.NamedHandler{Name: "S3MemBucketExists", Fn: getObjectBucketExists}
	req.Handlers.Send.PushBackNamed(bucketExists)
	getObject := aws.NamedHandler{Name: "S3MemGetObject", Fn: getObject}
	req.Handlers.Send.PushBackNamed(getObject)
	return s3.GetObjectRequest{Request: req, Input: input}
}

func getObjectBucketExists(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsBucketExist(req.Params.(*s3.GetObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.GetObjectInput).Bucket, nil, nil)
	}
}

func getObject(req *aws.Request) {
	if req.Error != nil {
		return
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
}
