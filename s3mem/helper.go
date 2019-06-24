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
	"testing"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.ibm.com/open-razee/s3mem-go/s3mem/s3memerr"
)

func NewDefaultS3MemService() *S3MemService {
	s3service := S3MemEndpointsID
	return NewS3MemService(s3service)
}

func NewTestS3MemService(t *testing.T) *S3MemService {
	return NewS3MemService(S3MemEndpointsID + "_" + strings.ToLower(t.Name()))
}

func NewS3MemService(s3service string) *S3MemService {
	if s3service == "" {
		s3service = S3MemEndpointsID
	}
	if _, ok := S3Store.S3MemServices[s3service]; !ok {
		S3Store.S3MemServices[s3service] = &S3MemService{
			Name:    s3service,
			Buckets: make(map[string]*Bucket, 0),
		}
		return S3Store.S3MemServices[s3service]
	}
	panic(fmt.Sprintf("The S3Store %s already exists", s3service))
}

func (s *S3MemService) DeleteDefaultS3MemService() {
	s.DeleteS3MemService(S3MemEndpointsID)
}

func (s *S3MemService) DeleteTestS3MemService(t *testing.T) {
	s.DeleteS3MemService(S3MemEndpointsID + "_" + strings.ToLower(t.Name()))
}

//Clear clears memory buckets and objects
func (s *S3MemService) DeleteS3MemService(s3service string) {
	delete(S3Store.S3MemServices, s3service)
}

func GetDefaultS3MemService() *S3MemService {
	if _, ok := S3Store.S3MemServices[S3MemEndpointsID]; !ok {
		panic(fmt.Sprintf("The S3Store %s doesn't not exist, please call NewDefaultS3MemService() before calling this function", S3MemEndpointsID))
	}
	return GetS3MemService(S3MemEndpointsID)
}

func GetTestS3MemService(t *testing.T) *S3MemService {
	return GetS3MemService(S3MemEndpointsID + "_" + strings.ToLower(t.Name()))
}

func GetS3MemService(s3service string) *S3MemService {
	if _, ok := S3Store.S3MemServices[s3service]; !ok {
		panic(fmt.Sprintf("The S3Store %s doesn't not exist, please call NewS3MemService(%s) before calling this function", s3service, s3service))
	}
	return S3Store.S3MemServices[s3service]
}

func (s *S3MemService) Lock() {
	s3memS3service := GetS3MemService(s.Name)
	s3memS3service.Mux.Lock()
}

func (s *S3MemService) Unlock() {
	s3memS3service := GetS3MemService(s.Name)
	s3memS3service.Mux.Unlock()
}

//GetBucket gets a bucket from memory
//The default s3service is S3MemEndpointsID
func (s *S3MemService) GetBucket(bucket *string) *s3.Bucket {
	s3memS3service := GetS3MemService(s.Name)
	if _, ok := s3memS3service.Buckets[*bucket]; !ok {
		return nil
	}
	return s3memS3service.Buckets[*bucket].Bucket
}

//IsBucketExist returns true if bucket exists
//The default s3service is S3MemEndpointsID
func (s *S3MemService) IsBucketExist(bucket *string) bool {
	s3memS3service := GetS3MemService(s.Name)
	_, ok := s3memS3service.Buckets[*bucket]
	return ok
}

//IsBucketEmpty returns true if bucket is empty
//The default s3service is S3MemEndpointsID
func (s *S3MemService) IsBucketEmpty(bucket *string) bool {
	s3memS3service := GetS3MemService(s.Name)
	return len(s3memS3service.Buckets[*bucket].Objects) == 0
}

//CreateBucket adds a bucket in memory
//The default s3service is S3MemEndpointsID
func (s *S3MemService) CreateBucket(b *s3.Bucket) {
	s3memS3service := GetS3MemService(s.Name)
	tc := time.Now()
	b.CreationDate = &tc
	s3memS3service.Buckets[*b.Name] = &Bucket{
		Bucket:  b,
		Objects: make(map[string]*VersionedObjects, 0),
	}
}

//DeleteBucket deletes an object from memory
//The default s3service is S3MemEndpointsID
func (s *S3MemService) DeleteBucket(bucket *string) {
	s3memS3service := GetS3MemService(s.Name)
	delete(s3memS3service.Buckets, *bucket)
}

//IsObjectExist returns true if object exists
//The default s3service is S3MemEndpointsID
func (s *S3MemService) IsObjectExist(bucket *string, key *string) bool {
	s3memS3service := GetS3MemService(s.Name)
	_, ok := s3memS3service.Buckets[*bucket].Objects[*key]
	return ok
}

//PutObject adds an object in memory return the object.
//The default s3service is S3MemEndpointsID
//Raise an error if a failure to read the body occurs
func (s *S3MemService) PutObject(bucket *string, key *string, body io.ReadSeeker) (*Object, *string, error) {
	s3memS3service := GetS3MemService(s.Name)
	if _, ok := s3memS3service.Buckets[*bucket]; !ok {
		s3memS3service.Buckets[*bucket].Objects = make(map[string]*VersionedObjects, 0)
	}
	if _, ok := s3memS3service.Buckets[*bucket].Objects[*key]; !ok {
		s3memS3service.Buckets[*bucket].Objects[*key] = &VersionedObjects{
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
	versioning := s3memS3service.Buckets[*bucket].VersioningConfiguration
	if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
		s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects, obj)
	} else {
		if len(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects) == 0 {
			s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects, obj)
		} else {
			s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[0] = obj
		}
	}
	versionId := strconv.Itoa(len(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects) - 1)
	return obj, &versionId, nil
}

//GetObject gets an object from memory return the Object and its versionID
//The default s3service is S3MemEndpointsID
//Raises an error the bucket or object doesn't exists or if the requested object is deleted,
func (s *S3MemService) GetObject(bucket *string, key *string, versionIDS *string) (object *Object, versionIDSOut *string, s3memerror s3memerr.S3MemError) {
	s3memS3service := GetS3MemService(s.Name)
	if _, ok := s3memS3service.Buckets[*bucket]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", nil, bucket, key, versionIDS)
	}
	if _, ok := s3memS3service.Buckets[*bucket].Objects[*key]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, versionIDS)
	}
	l := len(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects)
	if l > 0 {
		versioning := s3memS3service.Buckets[*bucket].VersioningConfiguration
		if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
			if versionIDS != nil {
				versionID, err := strconv.Atoi(*versionIDS)
				if err != nil {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "Version not a number", err, bucket, key, versionIDS)
				}
				if versionID >= l {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "", nil, bucket, key, versionIDS)
				}
				object = s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[versionID]
				versionIDSOut = versionIDS
				s3memerror = nil
			} else {
				object = s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
				versionID := strconv.Itoa(l - 1)
				versionIDSOut = &versionID
				s3memerror = nil
			}
		} else {
			object = s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
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
//The default s3service is S3MemEndpointsID
func (s *S3MemService) DeleteObject(bucket *string, key *string, versionIDS *string) (deleteMarkerOut *bool, deleteMarkerVersionIDOut *string, err s3memerr.S3MemError) {
	s3memS3service := GetS3MemService(s.Name)
	if _, ok := s3memS3service.Buckets[*bucket]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchBucket, "", err, bucket, key, versionIDS)
	}
	if _, ok := s3memS3service.Buckets[*bucket].Objects[*key]; !ok {
		return nil, nil, s3memerr.NewError(s3.ErrCodeNoSuchKey, "", nil, bucket, key, versionIDS)
	}
	l := len(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects)
	if l > 0 {
		versioning := s3memS3service.Buckets[*bucket].VersioningConfiguration
		if versioning != nil && versioning.Status == s3.BucketVersioningStatusEnabled {
			deleteMarker := true
			if versionIDS != nil {
				//if version provided then remove specific version
				versionID, err := strconv.Atoi(*versionIDS)
				if err != nil {
					return nil, nil, s3memerr.NewError(s3memerr.ErrCodeNoSuchVersion, "Version not a number", err, bucket, key, versionIDS)
				}
				if l-1 == versionID {
					s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects = s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[:l-1]
				} else {
					s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[:versionID], s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[versionID+1:]...)
				}
			} else {
				//if version not provided then add a marker object for the same version with no data
				deleteMarker = false
				currentVersionedObject := s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[l-1]
				versionID := strconv.Itoa(l - 1)
				deletedObject := &Object{
					DeletedObject: &s3.DeletedObject{
						DeleteMarker: &deleteMarker,
						Key:          currentVersionedObject.Object.Key,
						VersionId:    &versionID,
					},
				}
				s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects = append(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects, deletedObject)
				deleteMarkerVersionID := strconv.Itoa(len(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects) - 1)
				deleteMarkerVersionIDOut = &deleteMarkerVersionID
				deletedObject.DeletedObject.DeleteMarkerVersionId = deleteMarkerVersionIDOut
			}
			deleteMarkerOut = &deleteMarker
		} else {
			s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects = s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects[:l-1]
		}
	}
	if len(s3memS3service.Buckets[*bucket].Objects[*key].VersionedObjects) == 0 {
		delete(s3memS3service.Buckets[*bucket].Objects, *key)
	}
	return deleteMarkerOut, deleteMarkerVersionIDOut, nil
}

//PutBucketVersioning Sets the bucket in versionning mode
//The default s3service is S3MemEndpointsID
func (s *S3MemService) PutBucketVersioning(bucket *string, mfa *string, versioningConfiguration *s3.VersioningConfiguration) error {
	s3memS3service := GetS3MemService(s.Name)
	s3memS3service.Buckets[*bucket].MFA = mfa
	s3memS3service.Buckets[*bucket].VersioningConfiguration = versioningConfiguration
	return nil
}

//GetBucketVersioning gets the versioning configuration.
func (s *S3MemService) GetBucketVersioning(bucket *string) (*string, *s3.VersioningConfiguration) {
	s3memS3service := GetS3MemService(s.Name)
	if _, ok := s3memS3service.Buckets[*bucket]; !ok {
		return nil, nil
	}
	return s3memS3service.Buckets[*bucket].MFA, s3memS3service.Buckets[*bucket].VersioningConfiguration
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
