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
	"net/http"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

func (c *S3Mem) DeleteObjectRequest(input *s3.DeleteObjectInput) s3.DeleteObjectRequest {
	if input == nil {
		input = &s3.DeleteObjectInput{}
	}
	output := &s3.DeleteObjectOutput{}
	req := &aws.Request{
		Params:      input,
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	bucketExists := aws.NamedHandler{Name: "S3MemBucketExists", Fn: deleteObjectBucketExists}
	req.Handlers.Send.PushBackNamed(bucketExists)
	objectExists := aws.NamedHandler{Name: "S3MemObjectExists", Fn: deleteObjectObjectExists}
	req.Handlers.Send.PushBackNamed(objectExists)
	deleteObject := aws.NamedHandler{Name: "S3MemDeleteObject", Fn: deleteObject}
	req.Handlers.Send.PushBackNamed(deleteObject)
	return s3.DeleteObjectRequest{Request: req, Input: input}
}

func deleteObjectBucketExists(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsBucketExist(req.Params.(*s3.DeleteObjectInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteObjectInput).Bucket, nil, nil)
	}
}

func deleteObjectObjectExists(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsObjectExist(req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key, req.Params.(*s3.DeleteObjectInput).VersionId)
	}
}

func deleteObject(req *aws.Request) {
	if req.Error != nil {
		return
	}
	deleteMarker, versionID, err := DeleteObject(req.Params.(*s3.DeleteObjectInput).Bucket, req.Params.(*s3.DeleteObjectInput).Key, req.Params.(*s3.DeleteObjectInput).VersionId)
	if err != nil {
		req.Error = err
		return
	}
	req.Data.(*s3.DeleteObjectOutput).DeleteMarker = deleteMarker
	req.Data.(*s3.DeleteObjectOutput).VersionId = versionID
}
