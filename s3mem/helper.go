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
	"io"
	"io/ioutil"
	"log"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

//Clear clears memory buckets and objects
func Clear() {
	S3MemBuckets.Buckets = make(map[string]*Bucket, 0)
}

//GetBucket gets a bucket from memory
func GetBucket(bucket *string) *s3.Bucket {
	if _, ok := S3MemBuckets.Buckets[*bucket]; !ok {
		return nil
	}
	return S3MemBuckets.Buckets[*bucket].Bucket
}

func IsBucketExist(bucket *string) bool {
	_, ok := S3MemBuckets.Buckets[*bucket]
	return ok
}

func IsBucketEmpty(bucket *string) bool {
	return len(S3MemBuckets.Buckets[*bucket].Objects) == 0
}

//CreateBucket adds a bucket in memory
func CreateBucket(b *s3.Bucket) {
	S3MemBuckets.Buckets[*b.Name] = &Bucket{
		Bucket:  b,
		Objects: make(map[string]*VersionedObjects, 0),
	}
}

//DeleteBucket deletes an object from memory
func DeleteBucket(bucket *string) {
	delete(S3MemBuckets.Buckets, *bucket)
}

func IsObjectExist(bucket *string, key *string) bool {
	_, ok := S3MemBuckets.Buckets[*bucket].Objects[*key]
	return ok
}

//PutObject adds an object in memory
func PutObject(bucket *string, key *string, body io.ReadSeeker) (*Object, error) {
	if _, ok := S3MemBuckets.Buckets[*bucket]; !ok {
		S3MemBuckets.Buckets[*bucket].Objects = make(map[string]*VersionedObjects, 0)
	}
	if _, ok := S3MemBuckets.Buckets[*bucket].Objects[*key]; !ok {
		S3MemBuckets.Buckets[*bucket].Objects[*key] = &VersionedObjects{
			VersionedObjects: make([]*Object, 0),
		}
	}
	tc := time.Now()
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	obj := &Object{
		Object: &s3.Object{
			Key:          key,
			LastModified: &tc,
			StorageClass: "memory",
		},
		Content: content,
	}
	versioning := S3MemBuckets.Buckets[*bucket].VersioningConfiguration
	if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
		S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects, obj)
		// versionId := strconv.Itoa(len(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects - 1))
		// obj.VersionedObjects[versionId].VersionId = &versionId
	} else {
		S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects, obj)
	}
	return obj, nil
}

//GetObject gets an object from memory
func GetObject(bucket *string, key *string, versionIDS *string) (object *Object, versionIDSOut *string, s3memerror s3memerr.S3MemError) {
	if _, ok := S3MemBuckets.Buckets[*bucket]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, bucket, key, versionIDS)
	}
	if _, ok := S3MemBuckets.Buckets[*bucket].Objects[*key]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, versionIDS)
	}
	l := len(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects)
	log.Printf("Size %d:", l)
	if l > 0 {
		versioning := S3MemBuckets.Buckets[*bucket].VersioningConfiguration
		if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
			if versionIDS != nil {
				versionID, err := strconv.Atoi(*versionIDS)
				if err != nil {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "Version not a number", err, bucket, key, versionIDS)
				}
				if versionID >= l {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "", nil, bucket, key, versionIDS)
				}
				object = S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[versionID]
				versionIDSOut = versionIDS
				s3memerror = nil
			} else {
				object = S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
				versionID := strconv.Itoa(l - 1)
				versionIDSOut = &versionID
				s3memerror = nil
			}
		} else {
			object = S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
			versionID := strconv.Itoa(l - 1)
			versionIDSOut = &versionID
			s3memerror = nil
		}
		if object.DeletedObject != nil {
			return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "", nil, bucket, key, versionIDS)
		}
		log.Printf("object versionID %s", *versionIDSOut)
		return
	}
	return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, nil)
}

func DeleteObject(bucket *string, key *string, versionIDS *string) (deleteMarkerOut *bool, err s3memerr.S3MemError) {
	if _, ok := S3MemBuckets.Buckets[*bucket]; !ok {
		return nil, s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", err, bucket, key, versionIDS)
	}
	if _, ok := S3MemBuckets.Buckets[*bucket].Objects[*key]; !ok {
		return nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, versionIDS)
	}
	l := len(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects)
	if l > 0 {
		versioning := S3MemBuckets.Buckets[*bucket].VersioningConfiguration
		if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
			deleteMarker := true
			if versionIDS != nil {
				//if version provided then remove specific version
				versionID, err := strconv.Atoi(*versionIDS)
				if err != nil {
					return nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "Version not a number", err, bucket, key, versionIDS)
				}
				S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[:versionID], S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[versionID+1:]...)
			} else {
				//if version not provided then add a marker object for the same version with no data
				deleteMarker = false
				currentVersionedObject := S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
				versionID := strconv.Itoa(l - 1)
				deletedObject := &Object{
					DeletedObject: &s3.DeletedObject{
						DeleteMarker: &deleteMarker,
						Key:          currentVersionedObject.Object.Key,
						VersionId:    &versionID,
					},
				}
				S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects, deletedObject)
			}
			deleteMarkerOut = &deleteMarker
		} else {
			S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[:l-1]
		}
	}
	if len(S3MemBuckets.Buckets[*bucket].Objects[*key].VersionedObjects) == 0 {
		delete(S3MemBuckets.Buckets[*bucket].Objects, *key)
	}
	return deleteMarkerOut, nil
}

func PutBucketVersioning(bucket *string, mfa *string, versioningConfiguration *s3.VersioningConfiguration) error {
	S3MemBuckets.Buckets[*bucket].MFA = mfa
	S3MemBuckets.Buckets[*bucket].VersioningConfiguration = versioningConfiguration
	return nil
}

func GetBucketVersioning(bucket *string) (*string, *s3.VersioningConfiguration) {
	if _, ok := S3MemBuckets.Buckets[*bucket]; !ok {
		return nil, nil
	}
	return S3MemBuckets.Buckets[*bucket].MFA, S3MemBuckets.Buckets[*bucket].VersioningConfiguration
}
