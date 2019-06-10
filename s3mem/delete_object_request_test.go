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

func TestDeleteObjectRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	PutObject(&bucketName, &objectKey, strings.NewReader(string("test content")))
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)
	object, _, err := GetObject(&bucketName, &objectKey, nil)
	assert.Error(t, err)
	assert.Nil(t, object)
}

func TestDeleteObjectRequestBucketVersionedThenRestore(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Make bucket versioning
	PutBucketVersioning(&bucketName, nil, &s3.VersioningConfiguration{
		Status: s3.BucketVersioningStatusEnabled,
	})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	content := "test content"
	PutObject(&bucketName, &objectKey, strings.NewReader(content))
	//Request a client
	client := New(S3MemTestConfig)
	//Create the request
	req := client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	//Send the request
	deleteObjectOutput, err := req.Send(context.Background())
	assert.NoError(t, err)
	object, _, err := GetObject(&bucketName, &objectKey, nil)
	assert.Error(t, err)
	assert.Nil(t, object)

	//Restore object by delete marker
	req = client.DeleteObjectRequest(&s3.DeleteObjectInput{
		Bucket:    &bucketName,
		Key:       &objectKey,
		VersionId: deleteObjectOutput.VersionId,
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.NoError(t, err)

	object, _, err = GetObject(&bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object)

	assert.Equal(t, content, string(object.Content))

}
