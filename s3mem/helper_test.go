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

	"github.com/stretchr/testify/assert"
)

func TestParseObjectURL(t *testing.T) {
	url := "bucket/folder1/folder2/key"
	bucket, key, err := ParseObjectURL(&url)
	assert.NoError(t, err)
	assert.Equal(t, "bucket", *bucket)
	assert.Equal(t, "folder1/folder2/key", *key)
}
