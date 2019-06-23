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
)

func TestDeleteObjectRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	S3MemTestService.PutObject(&bucketName, &objectKey, strings.NewReader(string("test content")))
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)
	object, _, err := S3MemTestService.GetObject(&bucketName, &objectKey, nil)
	assert.Error(t, err)
	assert.Nil(t, object)
}

func TestDeleteObjectRequestBucketVersionedThenRestore(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	S3MemTestService.CreateBucket(&s3.Bucket{Name: &bucketName})
	//Make bucket versioning
	S3MemTestService.PutBucketVersioning(&bucketName, nil, &s3.VersioningConfiguration{
		Status: s3.BucketVersioningStatusEnabled,
	})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	content := "test content"
	S3MemTestService.PutObject(&bucketName, &objectKey, strings.NewReader(content))
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	//Send the request
	deleteObjectOutput, err := req.Send(context.Background())
	assert.NoError(t, err)
	object, _, err := S3MemTestService.GetObject(&bucketName, &objectKey, nil)
	assert.Error(t, err)
	assert.Nil(t, object)

	//Restore object by delete marker
	req = client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket:    &bucketName,
		Key:       &objectKey,
		VersionId: deleteObjectOutput.VersionId,
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.NoError(t, err)

	object, _, err = S3MemTestService.GetObject(&bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object)

	assert.Equal(t, content, string(object.Content))

}
