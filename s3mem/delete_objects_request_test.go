/*
################################################################################
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
*/

package s3mem

import (
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.com/IBM/s3mem-go/s3mem/s3memerr"
)

func TestDeleteObjectsRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey1 := "my-object1"
	S3MemTestService.PutObject(&bucketName, &objectKey1, strings.NewReader(string("test contents")))
	//Adding an Object directly in mem to prepare the test.
	objectKey2 := "my-object2"
	S3MemTestService.PutObject(&bucketName, &objectKey2, strings.NewReader(string("test contents")))
	//Request a client
	client := New(S3MemTestConfig)
	versionId := "1"
	//Create the request
	req := client.DeleteObjectsRequest(&s3.DeleteObjectsInput{
		Bucket: &bucketName,
		Delete: &s3.Delete{
			Objects: []s3.ObjectIdentifier{
				{Key: &objectKey1, VersionId: &versionId},
				{Key: &objectKey2, VersionId: &versionId},
			},
		},
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)
	object1, _, err := S3MemTestService.GetObject(&bucketName, &objectKey1, nil)
	assert.Error(t, err)
	assert.Nil(t, object1)
	object2, _, err := S3MemTestService.GetObject(&bucketName, &objectKey2, nil)
	assert.Error(t, err)
	assert.Nil(t, object2)
}

func TestDeleteObjectsRequestBucketNotExists(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey1 := "my-object1"
	S3MemTestService.PutObject(&bucketName, &objectKey1, strings.NewReader(string("test contents")))
	//Request a client
	client := New(S3MemTestConfig)
	versionId := "1"
	nonExistBucketName := strings.ToLower(t.Name()) + "-1"
	//Create the request
	req := client.DeleteObjectsRequest(&s3.DeleteObjectsInput{
		Bucket: &nonExistBucketName,
		Delete: &s3.Delete{
			Objects: []s3.ObjectIdentifier{
				{Key: &objectKey1, VersionId: &versionId},
			},
		},
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.Error(t, err)
	assert.Implements(t, (*s3memerr.S3MemError)(nil), err)
	assert.Equal(t, s3.ErrCodeNoSuchBucket, err.(s3memerr.S3MemError).Code())
}
