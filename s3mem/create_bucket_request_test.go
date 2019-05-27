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

func TestCreateBucketRequest(t *testing.T) {
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	//Create the request
	bucketName := strings.ToLower(t.Name())
	req := client.CreateBucketRequest(&s3.CreateBucketInput{
		Bucket: &bucketName,
	})
	//Send the request
	createBucketsOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	bucketGet := GetBucket(&bucketName)
	assert.NotNil(t, bucketGet)
	assert.Equal(t, bucketName, *bucketGet.Name)
	assert.Equal(t, bucketName, *createBucketsOutput.Location)
	assert.NotNil(t, bucketGet.CreationDate)
}
