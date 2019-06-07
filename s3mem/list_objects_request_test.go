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
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestListObjectssRequest(t *testing.T) {
	//Need to lock for testing as tests are running concurrently
	//and meanwhile another running test could change the stored buckets
	S3MemBuckets.Mux.Lock()
	defer S3MemBuckets.Mux.Unlock()

	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey1 := "1-my-object"
	content1 := "test content 1"
	PutObject(&bucketName, &objectKey1, strings.NewReader(string(content1)))
	objectKey2 := "2-my-object"
	content2 := "test content 2"
	PutObject(&bucketName, &objectKey2, strings.NewReader(string(content2)))

	//Request a client
	client := New()
	//Create the request
	req := client.ListObjectsRequest(&s3.ListObjectsInput{
		Bucket: &bucketName,
	})
	//Send the request
	listObjectsOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 2, len(listObjectsOutput.Contents))
	//Create the request
	prefix := "1"
	req = client.ListObjectsRequest(&s3.ListObjectsInput{
		Bucket: &bucketName,
		Prefix: &prefix,
	})
	//Send the request
	listObjectsOutput, err = req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 1, len(listObjectsOutput.Contents))
}
