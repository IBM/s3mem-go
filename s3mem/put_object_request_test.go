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
	AddBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	content := "test content"
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	//Create the request
	req := client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
		Body:   strings.NewReader(string(content)),
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.NoError(t, err)

	object := GetObject(&bucketName, &objectKey)
	assert.NotNil(t, object)

	assert.Equal(t, content, string(object.Content))
}
