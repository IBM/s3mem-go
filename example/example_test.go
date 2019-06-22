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
package example

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem"
)

var MyS3MemService *s3mem.S3MemService

func init() {
	MyS3MemService = s3mem.NewS3MemService(MyURLEndpoint)
}

func getTestConfig(S3MemService string) aws.Config {
	defaultResolver := endpoints.NewDefaultResolver()
	myCustomResolver := func(service, region string) (aws.Endpoint, error) {
		if service == s3.EndpointsID {
			return aws.Endpoint{
				URL: S3MemService,
			}, nil
		}
		return defaultResolver.ResolveEndpoint(service, region)
	}
	config := aws.Config{
		EndpointResolver: aws.EndpointResolverFunc(myCustomResolver),
	}
	return config
}

func TestGetObject(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucket := strings.ToLower(t.Name())
	MyS3MemService.CreateBucket(&s3.Bucket{Name: &bucket})
	//Adding an Object directly in mem to prepare the test.
	key := "my-object"
	content := "test content"
	MyS3MemService.PutObject(&bucket, &key, strings.NewReader(string(content)))
	//Request a client
	client := s3mem.New(getTestConfig(MyURLEndpoint))
	//Call the method to test
	b, err := GetObject(client, &bucket, &key)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, content, string(b))
}

//Here we use the default endpoint
func TestGetObjectWithDefaultEndpoint(t *testing.T) {
	//Request a client
	config := aws.Config{}
	client := s3mem.New(config)
	//Create default S3MemService
	defaultS3MemService := s3mem.NewDefaultS3MemService()
	//Adding bucket directly in mem to prepare the test.
	bucket := strings.ToLower(t.Name())
	// endpointResolver := config.EndpointResolver
	defaultS3MemService.CreateBucket(&s3.Bucket{Name: &bucket})
	//Adding an Object directly in mem to prepare the test.
	key := "my-object"
	content := "test content"
	defaultS3MemService.PutObject(&bucket, &key, strings.NewReader(string(content)))
	//Call the method to test
	b, err := GetObject(client, &bucket, &key)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, content, string(b))
	//Free up memory
	defaultS3MemService.DeleteDefaultS3MemService()
}

func TestGetObjectWithLog(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucket := strings.ToLower(t.Name())
	MyS3MemService.CreateBucket(&s3.Bucket{Name: &bucket})
	//Adding an Object directly in mem to prepare the test.
	key := "my-object"
	content := "test content"
	MyS3MemService.PutObject(&bucket, &key, strings.NewReader(string(content)))
	config := getTestConfig(MyURLEndpoint)
	config.LogLevel = aws.LogDebugWithHTTPBody
	config.Logger = aws.NewDefaultLogger()
	//Request a client
	client := s3mem.New(config)
	//Call the method to test
	b, err := GetObject(client, &bucket, &key)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, content, string(b))
}

func TestListBucketsRequest(t *testing.T) {
	//Creating a new S3MemService because at the end of this test
	//we want to test the bucket ordering and so we want a new
	//S3MemService to avoid interraction with other tests.
	S3MemService := strings.ToLower(t.Name())
	localS3MemService := s3mem.NewS3MemService(S3MemService)
	//Need to lock for testing as tests are running concurrently
	//and meanwhile another running test could change the stored buckets
	s3mem.S3Store.S3MemServices[S3MemService].Mux.Lock()
	defer s3mem.S3Store.S3MemServices[S3MemService].Mux.Unlock()
	//Adding bucket directly in mem to prepare the test.
	bucket0 := strings.ToLower(t.Name() + "0")
	bucket1 := strings.ToLower(t.Name() + "1")
	localS3MemService.CreateBucket(&s3.Bucket{Name: &bucket0})
	localS3MemService.CreateBucket(&s3.Bucket{Name: &bucket1})
	l := len(s3mem.S3Store.S3MemServices[S3MemService].Buckets)
	//Request a client
	client := s3mem.New(getTestConfig(S3MemService))
	//Call GetBuckets
	buckets, err := GetBuckets(client)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, l, len(buckets))
	//We can check each bucket name in that order as the
	//AWS S3 ListBucketRequest is supposed to return all bucket in
	//alphabetic order.
	assert.Equal(t, bucket0, buckets[0])
	assert.Equal(t, bucket1, buckets[1])
}

func TestListBucketsRequestWithLock(t *testing.T) {
	//Need to lock for testing as tests are running concurrently
	//and meanwhile another running test could change the stored buckets
	MyS3MemService.Lock()
	defer MyS3MemService.Unlock()
	//Adding bucket directly in mem to prepare the test.
	bucket0 := strings.ToLower(t.Name() + "0")
	bucket1 := strings.ToLower(t.Name() + "1")
	MyS3MemService.CreateBucket(&s3.Bucket{Name: &bucket0})
	MyS3MemService.CreateBucket(&s3.Bucket{Name: &bucket1})
	l := len(s3mem.S3Store.S3MemServices[MyURLEndpoint].Buckets)
	//Request a client
	client := s3mem.New(getTestConfig(MyURLEndpoint))
	//Call GetBuckets
	buckets, err := GetBuckets(client)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, l, len(buckets))
	//Here we can not test the order as the MyURLEndpoint
	//could get some extra buckets from other tests
}
