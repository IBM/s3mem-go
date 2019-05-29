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

func TestDeleteObjectsRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	AddBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey1 := "my-object1"
	AddObject(&bucketName, &objectKey1, strings.NewReader(string("test contents")))
	//Adding an Object directly in mem to prepare the test.
	objectKey2 := "my-object2"
	AddObject(&bucketName, &objectKey2, strings.NewReader(string("test contents")))
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
	object1 := GetObject(&bucketName, &objectKey1)
	assert.Nil(t, object1)
	object2 := GetObject(&bucketName, &objectKey2)
	assert.Nil(t, object2)
}
