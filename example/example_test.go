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
package example

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem"
)

func TestGetObject(t *testing.T) {
	//Adding bucket directly in mem to prepare the test.
	bucket := strings.ToLower(t.Name())
	s3mem.CreateBucket(&s3.Bucket{Name: &bucket})
	//Adding an Object directly in mem to prepare the test.
	key := "my-object"
	content := "test content"
	s3mem.PutObject(&bucket, &key, strings.NewReader(string(content)))
	//Request a client
	client := s3mem.New()
	//Call the method to test
	b, err := GetObject(client, &bucket, &key)
	//Assert the result
	assert.NoError(t, err)
	assert.Equal(t, content, string(b))
}
