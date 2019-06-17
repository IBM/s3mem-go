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
package example

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem"
)

func TestGetObject(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucket := strings.ToLower(t.Name())
	s3mem.CreateBucket(&s3.Bucket{Name: &bucket})
	//Adding an Object directly in mem to prepare the test.
	key := "my-object"
	content := "test content"
	s3mem.PutObject(&bucket, &key, strings.NewReader(string(content)))
	config := aws.Config{}
	//Request a client
	client := s3mem.New(config)
	//Call the method to test
	b, err := GetObject(client, &bucket, &key)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, content, string(b))
}

func TestGetObjectWithLog(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucket := strings.ToLower(t.Name())
	s3mem.CreateBucket(&s3.Bucket{Name: &bucket})
	//Adding an Object directly in mem to prepare the test.
	key := "my-object"
	content := "test content"
	s3mem.PutObject(&bucket, &key, strings.NewReader(string(content)))
	// config, err := external.LoadDefaultAWSConfig()
	// assert.NoError(t, err)
	config := aws.Config{
		LogLevel: aws.LogDebugWithHTTPBody,
		Logger:   aws.NewDefaultLogger(),
	}
	//Request a client
	client := s3mem.New(config)
	//Call the method to test
	b, err := GetObject(client, &bucket, &key)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, content, string(b))
}

func TestListBucketsRequest(t *testing.T) {
	//Need to lock for testing as tests are running concurrently
	//and meanwhile another running test could change the stored buckets
	s3mem.S3MemBuckets.Mux.Lock()
	defer s3mem.S3MemBuckets.Mux.Unlock()
	//Adding bucket directly in mem to prepare the test.
	bucket0 := strings.ToLower(t.Name() + "0")
	bucket1 := strings.ToLower(t.Name() + "1")
	s3mem.CreateBucket(&s3.Bucket{Name: &bucket0})
	s3mem.CreateBucket(&s3.Bucket{Name: &bucket1})
	l := len(s3mem.S3MemBuckets.Buckets)
	//Request a client
	config := aws.Config{}
	client := s3mem.New(config)
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
