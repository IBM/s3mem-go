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
	"testing"

	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/stretchr/testify/assert"
)

func TestNotImplemented(t *testing.T) {
	//Request a client
	client, err := NewClient()
	assert.NoError(t, err)
	assert.NotNil(t, client)
	input := &s3.AbortMultipartUploadInput{}
	req := client.AbortMultipartUploadRequest(input)
	assert.Error(t, req.Error)
	assert.Equal(t, s3memerr.ErrCodeNotImplemented, req.Error.(s3memerr.S3MemError).Code())
}
