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
)

func TestGetBucketVersioningRequest(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucketName := strings.ToLower(t.Name())
	CreateBucket(&s3.Bucket{Name: &bucketName})

	mfa := "122334 13445"

	//Add a VersionConfig
	PutBucketVersioning(&bucketName, &mfa, &s3.VersioningConfiguration{
		MFADelete: s3.MFADeleteEnabled,
		Status:    s3.BucketVersioningStatusEnabled,
	})
	//Request a client
	client := New(aws.Config{})

	//Create request
	req := client.GetBucketVersioningRequest(&s3.GetBucketVersioningInput{
		Bucket: &bucketName,
	})

	//Send the request
	getBucketVersioningOut, err := req.Send(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, s3.MFADeleteStatusEnabled, getBucketVersioningOut.MFADelete)
	assert.Equal(t, s3.BucketVersioningStatusEnabled, getBucketVersioningOut.Status)
}
