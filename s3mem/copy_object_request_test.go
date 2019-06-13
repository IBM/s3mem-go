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

func TestCopyObject(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})
	//Adding bucket directly in mem to prepare the test.
	bucketNameDest := strings.ToLower(t.Name() + "-dest")
	CreateBucket(&s3.Bucket{Name: &bucketNameDest})
	//Adding an Object directly in mem to prepare the test.
	objectKey := "my-object"
	content := "test content"
	source := bucketName + "/" + objectKey
	_, sourceVerionId, err := PutObject(&bucketName, &objectKey, strings.NewReader(string(content)))
	assert.NoError(t, err)
	//Request a client
	client := New(S3MemTestConfig)
	req := client.CopyObjectRequest(&s3.CopyObjectInput{
		Bucket:     &bucketNameDest,
		CopySource: &source,
	})
	objOut, err := req.Send(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, sourceVerionId, objOut.CopySourceVersionId)
	obj, _, err := GetObject(&bucketNameDest, &objectKey, nil)
	assert.NoError(t, err)
	assert.Equal(t, objectKey, *obj.Object.Key)
}
