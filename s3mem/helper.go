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
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

//Clear clears memory buckets and objects
func Clear(datastore string) {
	S3MemDatastores.Datastores[datastore].Buckets = make(map[string]*Bucket, 0)
}

func CreateDatastore(datastore string) *Buckets {
	datastore = GetDatastore(datastore)
	if _, ok := S3MemDatastores.Datastores[datastore]; !ok {
		S3MemDatastores.Datastores[datastore] = &Buckets{
			Buckets: make(map[string]*Bucket, 0),
		}
	}
	return S3MemDatastores.Datastores[datastore]
}

func GetDatastore(datastore string) string {
	if datastore == "" {
		return S3MemEndpointsID
	}
	return datastore
}

//GetBucket gets a bucket from memory
//The default datastore is S3MemEndpointsID
func GetBucket(datastore string, bucket *string) *s3.Bucket {
	s3memBuckets := CreateDatastore(datastore)
	if _, ok := s3memBuckets.Buckets[*bucket]; !ok {
		return nil
	}
	return s3memBuckets.Buckets[*bucket].Bucket
}

//IsBucketExist returns true if bucket exists
//The default datastore is S3MemEndpointsID
func IsBucketExist(datastore string, bucket *string) bool {
	s3memBuckets := CreateDatastore(datastore)
	_, ok := s3memBuckets.Buckets[*bucket]
	return ok
}

//IsBucketEmpty returns true if bucket is empty
//The default datastore is S3MemEndpointsID
func IsBucketEmpty(datastore string, bucket *string) bool {
	s3memBuckets := CreateDatastore(datastore)
	return len(s3memBuckets.Buckets[*bucket].Objects) == 0
}

//CreateBucket adds a bucket in memory
//The default datastore is S3MemEndpointsID
func CreateBucket(datastore string, b *s3.Bucket) {
	s3memBuckets := CreateDatastore(datastore)
	tc := time.Now()
	b.CreationDate = &tc
	s3memBuckets.Buckets[*b.Name] = &Bucket{
		Bucket:  b,
		Objects: make(map[string]*VersionedObjects, 0),
	}
}

//DeleteBucket deletes an object from memory
//The default datastore is S3MemEndpointsID
func DeleteBucket(datastore string, bucket *string) {
	s3memBuckets := CreateDatastore(datastore)
	delete(s3memBuckets.Buckets, *bucket)
}

//IsObjectExist returns true if object exists
//The default datastore is S3MemEndpointsID
func IsObjectExist(datastore string, bucket *string, key *string) bool {
	s3memBuckets := CreateDatastore(datastore)
	_, ok := s3memBuckets.Buckets[*bucket].Objects[*key]
	return ok
}

//PutObject adds an object in memory return the object.
//The default datastore is S3MemEndpointsID
//Raise an error if a failure to read the body occurs
func PutObject(datastore string, bucket *string, key *string, body io.ReadSeeker) (*Object, *string, error) {
	s3memBuckets := CreateDatastore(datastore)
	if _, ok := s3memBuckets.Buckets[*bucket]; !ok {
		s3memBuckets.Buckets[*bucket].Objects = make(map[string]*VersionedObjects, 0)
	}
	if _, ok := s3memBuckets.Buckets[*bucket].Objects[*key]; !ok {
		s3memBuckets.Buckets[*bucket].Objects[*key] = &VersionedObjects{
			VersionedObjects: make([]*Object, 0),
		}
	}
	tc := time.Now()
	content, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, nil, err
	}
	obj := &Object{
		Object: &s3.Object{
			Key:          key,
			LastModified: &tc,
			StorageClass: "memory",
		},
		Content: content,
	}
	versioning := s3memBuckets.Buckets[*bucket].VersioningConfiguration
	if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
		s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects, obj)
	} else {
		if len(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects) == 0 {
			s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects, obj)
		} else {
			s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[0] = obj
		}
	}
	versionId := strconv.Itoa(len(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects) - 1)
	return obj, &versionId, nil
}

//GetObject gets an object from memory return the Object and its versionID
//The default datastore is S3MemEndpointsID
//Raises an error the bucket or object doesn't exists or if the requested object is deleted,
func GetObject(datastore string, bucket *string, key *string, versionIDS *string) (object *Object, versionIDSOut *string, s3memerror s3memerr.S3MemError) {
	s3memBuckets := CreateDatastore(datastore)
	if _, ok := s3memBuckets.Buckets[*bucket]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, bucket, key, versionIDS)
	}
	if _, ok := s3memBuckets.Buckets[*bucket].Objects[*key]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, versionIDS)
	}
	l := len(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects)
	if l > 0 {
		versioning := s3memBuckets.Buckets[*bucket].VersioningConfiguration
		if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
			if versionIDS != nil {
				versionID, err := strconv.Atoi(*versionIDS)
				if err != nil {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "Version not a number", err, bucket, key, versionIDS)
				}
				if versionID >= l {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "", nil, bucket, key, versionIDS)
				}
				object = s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[versionID]
				versionIDSOut = versionIDS
				s3memerror = nil
			} else {
				object = s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
				versionID := strconv.Itoa(l - 1)
				versionIDSOut = &versionID
				s3memerror = nil
			}
		} else {
			object = s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
			versionID := strconv.Itoa(l - 1)
			versionIDSOut = &versionID
			s3memerror = nil
		}
		if object.DeletedObject != nil {
			return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "", nil, bucket, key, versionIDS)
		}
		return
	}
	return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, nil)
}

//DeleteObject Deletes an object
//The default datastore is S3MemEndpointsID
func DeleteObject(datastore string, bucket *string, key *string, versionIDS *string) (deleteMarkerOut *bool, deleteMarkerVersionIDOut *string, err s3memerr.S3MemError) {
	s3memBuckets := CreateDatastore(datastore)
	if _, ok := s3memBuckets.Buckets[*bucket]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", err, bucket, key, versionIDS)
	}
	if _, ok := s3memBuckets.Buckets[*bucket].Objects[*key]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, versionIDS)
	}
	l := len(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects)
	if l > 0 {
		versioning := s3memBuckets.Buckets[*bucket].VersioningConfiguration
		if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
			deleteMarker := true
			if versionIDS != nil {
				//if version provided then remove specific version
				versionID, err := strconv.Atoi(*versionIDS)
				if err != nil {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "Version not a number", err, bucket, key, versionIDS)
				}
				if l-1 == versionID {
					s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[:l-1]
				} else {
					s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[:versionID], s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[versionID+1:]...)
				}
			} else {
				//if version not provided then add a marker object for the same version with no data
				deleteMarker = false
				currentVersionedObject := s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
				versionID := strconv.Itoa(l - 1)
				deletedObject := &Object{
					DeletedObject: &s3.DeletedObject{
						DeleteMarker: &deleteMarker,
						Key:          currentVersionedObject.Object.Key,
						VersionId:    &versionID,
					},
				}
				s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects, deletedObject)
				deleteMarkerVersionID := strconv.Itoa(len(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects) - 1)
				deleteMarkerVersionIDOut = &deleteMarkerVersionID
				deletedObject.DeletedObject.DeleteMarkerVersionId = deleteMarkerVersionIDOut
			}
			deleteMarkerOut = &deleteMarker
		} else {
			s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects = s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects[:l-1]
		}
	}
	if len(s3memBuckets.Buckets[*bucket].Objects[*key].VersionedObjects) == 0 {
		delete(s3memBuckets.Buckets[*bucket].Objects, *key)
	}
	return deleteMarkerOut, deleteMarkerVersionIDOut, nil
}

//PutBucketVersioning Sets the bucket in versionning mode
//The default datastore is S3MemEndpointsID
func PutBucketVersioning(datastore string, bucket *string, mfa *string, versioningConfiguration *s3.VersioningConfiguration) error {
	s3memBuckets := CreateDatastore(datastore)
	s3memBuckets.Buckets[*bucket].MFA = mfa
	s3memBuckets.Buckets[*bucket].VersioningConfiguration = versioningConfiguration
	return nil
}

//GetBucketVersioning gets the versioning configuration.
func GetBucketVersioning(datastore string, bucket *string) (*string, *s3.VersioningConfiguration) {
	s3memBuckets := CreateDatastore(datastore)
	if _, ok := s3memBuckets.Buckets[*bucket]; !ok {
		return nil, nil
	}
	return s3memBuckets.Buckets[*bucket].MFA, s3memBuckets.Buckets[*bucket].VersioningConfiguration
}

func CreateUser(canonicalID, email *string) error {
	if _, ok := S3MemUsers.Users[*canonicalID]; ok {
		return fmt.Errorf("User %s already exists", *canonicalID)
	}
	S3MemUsers.Users[*canonicalID] = &User{
		CanonicalID: *canonicalID,
		Email:       *email,
	}
	return nil
}

func GetUser(canonicalID, email *string) (*User, error) {
	if canonicalID == nil {
		user, err := searchUserByEmail(email)
		if err != nil {
			return nil, err
		}
		return user, nil
	}
	if user, ok := S3MemUsers.Users[*canonicalID]; ok {
		return user, nil
	}
	return nil, fmt.Errorf("User with email %s not found", *email)
}

func searchUserByEmail(email *string) (*User, error) {
	var user *User
	for _, v := range S3MemUsers.Users {
		if v.Email == *email {
			user = v
			break
		}
	}
	if user == nil {
		return nil, fmt.Errorf("User with email %s not found", *email)
	}
	return user, nil
}

func ParseObjectURL(url *string) (bucket, key *string, err error) {
	segs := strings.SplitN(*url, "/", 2)
	if len(segs) < 2 {
		return nil, nil, errors.New("Malformed url")
	}
	return &segs[0], &segs[1], nil
}
