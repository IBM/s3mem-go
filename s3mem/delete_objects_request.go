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

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const opDeleteObjects = "DeleteObjects"

//DeleteObjectsRequest ...
func (c *Client) DeleteObjectsRequest(input *s3.DeleteObjectsInput) s3.DeleteObjectsRequest {
	if input == nil {
		input = &s3.DeleteObjectsInput{}
	}
	output := &s3.DeleteObjectsOutput{
		Deleted: make([]s3.DeletedObject, 0),
		Errors:  make([]s3.Error, 0),
	}
	op := &aws.Operation{
		Name:       opDeleteObjects,
		HTTPMethod: "POST",
		HTTPPath:   "/{Bucket}?delete",
	}
	req := c.NewRequest(op, input, output)
	return s3.DeleteObjectsRequest{Request: req, Input: input}
}

func deleteObjects(req *aws.Request) {
	if !IsBucketExist(req.Params.(*s3.DeleteObjectsInput).Bucket) {
		req.Error = s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, req.Params.(*s3.DeleteObjectsInput).Bucket, nil, nil)
		return
	}
	for _, obj := range req.Params.(*s3.DeleteObjectsInput).Delete.Objects {
		deleteMarker, versionID, err := DeleteObject(req.Params.(*s3.DeleteObjectsInput).Bucket, obj.Key, obj.VersionId)
		if err != nil {
			req.Data.(*s3.DeleteObjectsOutput).Errors = append(req.Data.(*s3.DeleteObjectsOutput).Errors, err.Convert2S3Error(obj.Key, obj.VersionId))
		}
		req.Data.(*s3.DeleteObjectsOutput).Deleted = append(req.Data.(*s3.DeleteObjectsOutput).Deleted, s3.DeletedObject{
			DeleteMarker:          deleteMarker,
			DeleteMarkerVersionId: versionID,
			VersionId:             obj.VersionId,
			Key:                   obj.Key,
		})
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.DeleteObjectsOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
