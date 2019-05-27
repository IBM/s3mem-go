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
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestGetObjectRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	AddBucket(&s3.Bucket{Name: &bucketName})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	content := "test content"
	AddObject(&bucketName, &objectKey, strings.NewReader(string(content)))
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	//Create the request
	req := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: &bucketName,
		Key:    &objectKey,
	})
	//Send the request
	object, err := req.Send(context.Background())
	assert.NoError(t, err)

	assert.NotNil(t, object.Body)

	buf := new(bytes.Buffer)
	buf.ReadFrom(object.Body)
	newBytes := buf.Bytes()
	assert.Equal(t, content, string(newBytes))
}
