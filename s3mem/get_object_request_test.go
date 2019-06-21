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
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestGetObjectRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(S3MemEndpointsID, &s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	content := "test content"
	PutObject(S3MemEndpointsID, &bucketName, &objectKey, strings.NewReader(string(content)))
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	//Send the request
	object, err := req.Send(context.Background())
	assert.NoError(t, err)

	assert.NotNil(t, object.Body)

	buf := new(bytes.Buffer)
	buf.ReadFrom(object.Body)
	newBytes := buf.Bytes()
	assert.Equal(t, content, string(newBytes))
}

func TestGetObjectRequestWithVersioningBucket(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(S3MemEndpointsID, &s3.Bucket{Name: &bucketName})
	//Make bucket versioning
	PutBucketVersioning(S3MemEndpointsID, &bucketName, nil, &s3.VersioningConfiguration{
		Status: s3.BucketVersioningStatusEnabled,
	})
	//Adding an Object
	objectKey := "1-my-object"
	content1 := "test content 1"
	PutObject(S3MemEndpointsID, &bucketName, &objectKey, strings.NewReader(string(content1)))
	content2 := "test content 2"
	PutObject(S3MemEndpointsID, &bucketName, &objectKey, strings.NewReader(string(content2)))
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request to get the last version
	req := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	//Send the request
	object, err := req.Send(context.Background())
	assert.NoError(t, err)

	assert.NotNil(t, object.Body)

	buf := new(bytes.Buffer)
	buf.ReadFrom(object.Body)
	newBytes := buf.Bytes()
	assert.Equal(t, content2, string(newBytes))

	assert.Equal(t, "1", *object.VersionId)

	//Create the request a specific version
	versionIDS := "0"
	req = client.GetObjectRequest(&s3.GetObjectInput{
		Bucket:    &bucketName,
		Key:       &objectKey,
		VersionId: &versionIDS,
	})
	//Send the request
	object, err = req.Send(context.Background())
	assert.NoError(t, err)

	assert.NotNil(t, object.Body)

	buf = new(bytes.Buffer)
	buf.ReadFrom(object.Body)
	newBytes = buf.Bytes()
	assert.Equal(t, content1, string(newBytes))

	assert.Equal(t, "0", *object.VersionId)
}
