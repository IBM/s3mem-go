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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var checkConfigHandler = aws.NamedHandler{Name: "s3mem.checkConfig", Fn: checkConfig}
var checkACLHandler = aws.NamedHandler{Name: "s3mem.checkACL", Fn: checkACL}
var sendHandler = aws.NamedHandler{Name: "s3mem.sendHandler", Fn: send}

func checkConfig(r *aws.Request) {
}

func checkACL(r *aws.Request) {
	switch r.Operation.Name {
	case "CreateBucketRequest":
	}
}

func send(r *aws.Request) {
	switch r.Params.(type) {
	case *s3.CopyObjectInput:
		copyObject(r)
	case *s3.CreateBucketInput:
		createBucket(r)
	case *s3.DeleteBucketInput:
		deleteBucket(r)
	case *s3.DeleteObjectInput:
		deleteObject(r)
	case *s3.DeleteObjectsInput:
		deleteObjects(r)
	case *s3.GetBucketVersioningInput:
		getBucketVersioning(r)
	case *s3.GetObjectInput:
		getObject(r)
	case *s3.ListBucketsInput:
		listBuckets(r)
	case *s3.ListObjectsInput:
		listObjects(r)
	case *s3.PutBucketVersioningInput:
		putBucketVersioning(r)
	case *s3.PutObjectInput:
		putObject(r)
	}
}
