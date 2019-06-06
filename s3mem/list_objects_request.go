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
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

func (c *S3Mem) ListObjectsRequest(input *s3.ListObjectsInput) s3.ListObjectsRequest {
	if input == nil {
		input = &s3.ListObjectsInput{}
	}
	output := &s3.ListObjectsOutput{}
	operation := &aws.Operation{}
	req := &aws.Request{
		Params:      input,
		Data:        output,
		Operation:   operation,
		HTTPRequest: &http.Request{},
	}
	bucketExists := aws.NamedHandler{Name: "S3MemBucketExists", Fn: listObjectsBucketExists}
	req.Handlers.Send.PushBackNamed(bucketExists)
	listObjects := aws.NamedHandler{Name: "S3MemListObjects", Fn: listObjects}
	req.Handlers.Send.PushBackNamed(listObjects)
	return s3.ListObjectsRequest{Request: req, Input: input, Copy: c.ListObjectsRequest}
}

func listObjectsBucketExists(req *aws.Request) {
	if req.Error != nil {
		return
	}
	if !IsBucketExist(req.Params.(*s3.ListObjectsInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.ListObjectsInput).Bucket, nil, nil)
	}
}

func listObjects(req *aws.Request) {
	if req.Error != nil {
		return
	}
	req.Data.(*s3.ListObjectsOutput).Contents = make([]s3.Object, 0)
	bucket := req.Params.(*s3.ListObjectsInput).Bucket
	prefix := req.Params.(*s3.ListObjectsInput).Prefix
	for _, obj := range S3MemBuckets.Buckets[*bucket].Objects {
		if prefix != nil {
			if strings.HasPrefix(*obj.VersionedObjects[0].Object.Key, *prefix) {
				req.Data.(*s3.ListObjectsOutput).Contents = append(req.Data.(*s3.ListObjectsOutput).Contents, *obj.VersionedObjects[0].Object)
			}
		} else {
			req.Data.(*s3.ListObjectsOutput).Contents = append(req.Data.(*s3.ListObjectsOutput).Contents, *obj.VersionedObjects[0].Object)
		}
	}
}
