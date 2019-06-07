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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

func TestDeleteBucketRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Request a client
	client := New(aws.Config{})
	//Create the request
	req := client.DeleteBucketRequest(&s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)
	bucketGet := GetBucket(&bucketName)
	assert.Nil(t, bucketGet)
}

func TestDeleteNotEmptyBucket(t *testing.T) {
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	PutObject(&bucketName, &objectKey, strings.NewReader(string("test content")))
	//Request a client
	client := New(aws.Config{})
	//Create the request
	req := client.DeleteBucketRequest(&s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.Error(t, err)
	assert.Implements(t, (*s3memerr.S3MemError)(nil), err)
	assert.Equal(t, s3memerr.ErrCodeBucketNotEmpty, err.(s3memerr.S3MemError).Code())
}
