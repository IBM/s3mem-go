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

func TestDeleteBucketRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	AddBucket(&s3.Bucket{Name: &bucketName})
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	//Create the request
	req := client.DeleteBucketRequest(&s3.DeleteBucketInput{
		Bucket: &bucketName,
	})
	//Send the request
	_, err = req.Send(context.Background())
	assert.NoError(t, err)
	bucketGet := GetBucket(&bucketName)
	assert.Nil(t, bucketGet)
}
