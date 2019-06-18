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

func TestCopyObject(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding bucket directly in mem to prepare the test.
	bucketNameDest := strings.ToLower(t.Name() + "-dest")
	CreateBucket(&s3.Bucket{Name: &bucketNameDest})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	content := "test content"
	source := bucketName + "/" + objectKey
	_, sourceVerionId, err := PutObject(&bucketName, &objectKey, strings.NewReader(string(content)))
	assert.NoError(t, err)
	//Request a client
	client := New(S3MemTestConfig)
	req := client.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:     &bucketNameDest,
		CopySource: &source,
	})
	objOut, err := req.Send(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, sourceVerionId, objOut.CopySourceVersionId)
	obj, _, err := GetObject(&bucketNameDest, &objectKey, nil)
	assert.NoError(t, err)
	assert.Equal(t, objectKey, *obj.Object.Key)
}
