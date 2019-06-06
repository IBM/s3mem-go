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
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

func TestDeleteObjectsRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey1 := "my-object1"
	PutObject(&bucketName, &objectKey1, strings.NewReader(string("test contents")))
	//Adding an Object directly in mem to prepare the test.
	objectKey2 := "my-object2"
	PutObject(&bucketName, &objectKey2, strings.NewReader(string("test contents")))
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	versionId := "1"
	//Create the request
	req := client.DeleteObjectsRequest(&s3.DeleteObjectsInput{
		Bucket: &bucketName,
		Delete: &s3.Delete{
			Objects: []s3.ObjectIdentifier{
				{Key: &objectKey1, VersionId: &versionId},
				{Key: &objectKey2, VersionId: &versionId},
			},
		},
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.NoError(t, err)
	object1, _, err := GetObject(&bucketName, &objectKey1, nil)
	assert.Error(t, err)
	assert.Nil(t, object1)
	object2, _, err := GetObject(&bucketName, &objectKey2, nil)
	assert.Error(t, err)
	assert.Nil(t, object2)
}

func TestDeleteObjectsRequestBucketNotExists(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey1 := "my-object1"
	PutObject(&bucketName, &objectKey1, strings.NewReader(string("test contents")))
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	versionId := "1"
	nonExistBucketName := strings.ToLower(t.Name()) + "-1"
	//Create the request
	req := client.DeleteObjectsRequest(&s3.DeleteObjectsInput{
		Bucket: &nonExistBucketName,
		Delete: &s3.Delete{
			Objects: []s3.ObjectIdentifier{
				{Key: &objectKey1, VersionId: &versionId},
			},
		},
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.Error(t, err)
	assert.Implements(t, (*s3memerr.S3MemError)(nil), err)
	assert.Equal(t, s3.ErrCodeNoSuchBucket, err.(s3memerr.S3MemError).Code())
}
