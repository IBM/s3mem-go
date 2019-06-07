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

func TestPutObjectRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object
	objectKey := "my-object"
	content := "test content"
	//Request a client
	client := New()
	//Create the request
	req := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   strings.NewReader(string(content)),
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)

	object, _, err := GetObject(&bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object)

	assert.Equal(t, content, string(object.Content))
}

func TestPutObjectRequestWithVersioningBucket(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Make bucket versioning
	PutBucketVersioning(&bucketName, nil, &s3.VersioningConfiguration{
		Status: s3.BucketVersioningStatusEnabled,
	})
	//Adding an Object
	objectKey := "my-object-1"
	content1 := "test content 1"
	//Request a client
	client := New()
	//Create the request
	req := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   strings.NewReader(string(content1)),
	})
	//Send the request
	_, err := req.Send(context.Background())
	assert.NoError(t, err)

	object1, _, err := GetObject(&bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object1)

	assert.Equal(t, content1, string(object1.Content))

	content2 := "test content 2"

	//Create the request
	req = client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   strings.NewReader(string(content2)),
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.NoError(t, err)

	//Get last version
	object2, versionID, err := GetObject(&bucketName, &objectKey, nil)
	assert.NoError(t, err)
	assert.NotNil(t, object2)
	assert.Equal(t, "1", *versionID)

	assert.Equal(t, content2, string(object2.Content))

	//Get Specific version
	versionIDS := "0"
	object3, versionID, err := GetObject(&bucketName, &objectKey, &versionIDS)
	assert.NoError(t, err)
	assert.NotNil(t, object3)
	assert.Equal(t, content1, string(object1.Content))
	assert.Equal(t, "0", *versionID)

}
