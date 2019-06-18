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
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/stretchr/testify/assert"
)

func TestClientLoggerWrite(t *testing.T) {
	logger := &logWriter{
		Logger: aws.NewDefaultLogger(),
		buf:    new(bytes.Buffer),
	}
	l := logger.buf.Len()
	n, err := logger.Write([]byte("b"))
	assert.NoError(t, err)
	assert.Equal(t, l+1, n)
}

// func TestClientLoggerClose(t *testing.T) {
// 	logger := &logWriter{
// 		Logger: aws.NewDefaultLogger(),
// 		buf:    new(bytes.Buffer),
// 	}
// 	err := logger.
// }
