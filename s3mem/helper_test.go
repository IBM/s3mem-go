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
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestGetDefaultS3MemServiceNotExists(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	GetS3MemService(strings.ToLower(t.Name()))
	assert.Fail(t, "Panic was expected as the S3Store doesn't exist yet")
}

func TestNewS3MemServiceAlreadyExists(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	NewS3MemService(strings.ToLower(t.Name()))
	NewS3MemService(strings.ToLower(t.Name()))
	assert.Fail(t, "Panic was expected as the S3Store doesn't exist yet")
}

func TestDeleteTestSevice(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	s := NewTestS3MemService(t)
	bucketName := strings.ToLower(t.Name())
	s.CreateBucket(&s3.Bucket{Name: &bucketName})
	s.DeleteTestS3MemService(t)
	GetTestS3MemService(t)
	assert.Fail(t, "Panic was expected as the S3Store doesn't exist anymore")
}
func TestParseObjectURL(t *testing.T) {
	url := "bucket/folder1/folder2/key"
	bucket, key, err := ParseObjectURL(&url)
	assert.NoError(t, err)
	assert.Equal(t, "bucket", *bucket)
	assert.Equal(t, "folder1/folder2/key", *key)
}
