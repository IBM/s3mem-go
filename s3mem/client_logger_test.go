/*
################################################################################
# Copyright 2019 IBM Corp. All Rights Reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
################################################################################
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
