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

func TestPutObjectRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(S3MemEndpointsID, &s3.Bucket{Name: &bucketName})
	//Adding an Object
	objectKey := "my-object"
	content := "test content"
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   strings.NewReader(string(content)),
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)

	object, _, err := GetObject(S3MemEndpointsID, &bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object)

	assert.Equal(t, content, string(object.Content))
}

func TestPutObjectRequestWithVersioningBucket(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(S3MemEndpointsID, &s3.Bucket{Name: &bucketName})
	//Make bucket versioning
	PutBucketVersioning(S3MemEndpointsID, &bucketName, nil, &s3.VersioningConfiguration{
		Status: s3.BucketVersioningStatusEnabled,
	})
	//Adding an Object
	objectKey := "my-object-1"
	content1 := "test content 1"
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   strings.NewReader(string(content1)),
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)

	object1, _, err := GetObject(S3MemEndpointsID, &bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object1)

	assert.Equal(t, content1, string(object1.Content))

	content2 := "test content 2"

	//Create the request
	req = client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   strings.NewReader(string(content2)),
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.NoError(t, err)

	//Get last version
	object2, versionID, err := GetObject(S3MemEndpointsID, &bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object2)
	assert.Equal(t, "1", *versionID)

	assert.Equal(t, content2, string(object2.Content))

	//Get Specific version
	versionIDS := "0"
	object3, versionID, err := GetObject(S3MemEndpointsID, &bucketName, &objectKey, &versionIDS)
	assert.NoError(t, err)
	assert.NotNil(t, object3)
	assert.Equal(t, content1, string(object1.Content))
	assert.Equal(t, "0", *versionID)

}
