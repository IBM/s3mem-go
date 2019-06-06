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

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *S3Mem) DeleteObjectsRequest(input *s3.DeleteObjectsInput) s3.DeleteObjectsRequest {
	if input == nil {
		input = &s3.DeleteObjectsInput{}
	}
	output := &s3.DeleteObjectsOutput{
		Deleted: make([]s3.DeletedObject, 0),
		Errors:  make([]s3.Error, 0),
	}
	req := &aws.Request{
		Params:      input,
		Data:        output,
		HTTPRequest: &http.Request{},
	}
	bucketExists := aws.NamedHandler{Name: "S3MemBucketExists", Fn: deleteObjectsBucketExists}
	req.Handlers.Send.PushBackNamed(bucketExists)
	deleteObjects := aws.NamedHandler{Name: "S3MemDeleteObjects", Fn: deleteObjects}
	req.Handlers.Send.PushBackNamed(deleteObjects)
	return s3.DeleteObjectsRequest{Request: req, Input: input, Copy: c.DeleteObjectsRequest}
}

func deleteObjectsBucketExists(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsBucketExist(req.Params.(*s3.DeleteObjectsInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteObjectsInput).Bucket, nil, nil)
	}
}

func deleteObjects(req *aws.Request) {
	if req.Error != nil {
		return
	}
	for _, obj := range req.Params.(*s3.DeleteObjectsInput).Delete.Objects {
		deleteMarker, err := DeleteObject(req.Params.(*s3.DeleteObjectsInput).Bucket, obj.Key, obj.VersionId)
		if err != nil {
			req.Data.(*s3.DeleteObjectsOutput).Errors = append(req.Data.(*s3.DeleteObjectsOutput).Errors, err.Convert2S3Error(obj.Key, obj.VersionId))
		}
		req.Data.(*s3.DeleteObjectsOutput).Deleted = append(req.Data.(*s3.DeleteObjectsOutput).Deleted, s3.DeletedObject{
			DeleteMarker:          deleteMarker,
			DeleteMarkerVersionId: obj.VersionId,
			VersionId:             obj.VersionId,
			Key:                   obj.Key,
		})
	}
}
