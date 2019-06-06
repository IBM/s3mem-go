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

package s3memerr

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const (
	ErrCodeBucketNotEmpty                          = "BucketNotEmpty"
	ErrCodeBucketAlreadyExists                     = "BucketAlreadyExists"
	ErrCodeIllegalVersioningConfigurationException = "IllegalVersioningConfigurationException"
	ErrCodeNoSuchVersion                           = "ErrCodeNoSuchVersion"
)

// Factory interface for creating Error instances.
type Factory interface {
	NewError(code, message string, origErr error, bucket, key, versionId *string) S3MemError
}

type errorFactory struct{}

// NewFactory Create a new default Factory for creating S3 instances
func NewFactory() Factory {
	return &errorFactory{}
}

type S3MemError interface {
	// Satisfy the generic error interface.
	awserr.Error

	Bucket() *string

	Key() *string

	VersionId() *string

	Convert2S3Error(key, versionId *string) s3.Error
}

type s3memError struct {
	code      string
	message   string
	origErr   error
	bucket    *string
	key       *string
	versionId *string
}

func NewError(code, message string, origErr error, bucket, key, versionId *string) S3MemError {
	errf := NewFactory()
	return errf.NewError(code, message, origErr, bucket, key, versionId)
}

func (errf errorFactory) NewError(code, message string, origErr error, bucket, key, versionId *string) S3MemError {
	return &s3memError{
		code:    code,
		message: message,
		origErr: origErr,
	}
}

func (s3memError s3memError) Code() string {
	return s3memError.code
}

func (s3memError s3memError) Message() string {
	return s3memError.message
}
func (s3memError s3memError) OrigErr() error {
	return s3memError.origErr
}

func (s3memError s3memError) Bucket() *string {
	return s3memError.bucket
}

func (s3memError s3memError) Key() *string {
	return s3memError.key
}

func (s3memError s3memError) VersionId() *string {
	return s3memError.versionId
}

func (s3memError s3memError) Error() string {
	var errMsg string
	if s3memError.origErr != nil {
		errMsg = s3memError.origErr.Error()
	}
	var bucket string
	if s3memError.bucket != nil {
		bucket = *s3memError.bucket
	}
	var key string
	if s3memError.key != nil {
		key = *s3memError.key
	}
	var versionId string
	if s3memError.versionId != nil {
		key = *s3memError.versionId
	}

	return fmt.Sprintf("Code: %s, Message: %s, error: %s, bucket: %s, key: %s, versionId: %s", s3memError.code, s3memError.message, errMsg, bucket, key, versionId)
}

func (s3memError s3memError) Convert2S3Error(key, versionId *string) s3.Error {
	code := s3memError.Code()
	message := s3memError.Message()
	return s3.Error{
		Code:      &code,
		Message:   &message,
		Key:       key,
		VersionId: versionId,
	}
}
