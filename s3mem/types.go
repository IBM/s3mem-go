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
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type S3Mem struct {
}

type Buckets struct {
	Buckets map[string]*Bucket
	Mux     sync.Mutex
}

//(bucket/key/version)
type Bucket struct {
	Bucket                  *s3.Bucket
	MFA                     *string
	VersioningConfiguration *s3.VersioningConfiguration
	Objects                 map[string]*VersionedObjects
	Mux                     sync.Mutex
}

type VersionedObjects struct {
	VersionedObjects []*Object
}

type Object struct {
	Object        *s3.Object
	DeletedObject *s3.DeletedObject
	Content       []byte
}
