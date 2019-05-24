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

func TestListBucketsRequest(t *testing.T) {
	//Need to lock for testing as tests are running concurrently
	//and meanwhile another running test could change the stored buckets
	S3MemBuckets.Mux.Lock()
	defer S3MemBuckets.Mux.Unlock()

	//Adding bucket directly in mem to prepare the test.
	bucket0 := strings.ToLower(t.Name() + "0")
	bucket1 := strings.ToLower(t.Name() + "1")
	AddBucket(&s3.Bucket{Name: &bucket0})
	AddBucket(&s3.Bucket{Name: &bucket1})
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	//Create the request
	req := client.ListBucketsRequest(&s3.ListBucketsInput{})
	//Send the request
	listBucketsOutput, err := req.Send(context.Background())
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, 2, len(listBucketsOutput.Buckets))
}
