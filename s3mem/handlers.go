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
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var checkConfigHandler = aws.NamedHandler{Name: "s3mem.checkConfig", Fn: checkConfig}
var checkACLHandler = aws.NamedHandler{Name: "s3mem.checkACL", Fn: checkACL}
var sendHandler = aws.NamedHandler{Name: "s3mem.sendHandler", Fn: send}

func checkConfig(r *aws.Request) {
}

func checkACL(r *aws.Request) {
	switch r.Operation.Name {
	case "CreateBucketRequest":
	}
}

func send(r *aws.Request) {
	switch r.Params.(type) {
	case *s3.CopyObjectInput:
		copyObject(r)
	case *s3.CreateBucketInput:
		createBucket(r)
	case *s3.DeleteBucketInput:
		deleteBucket(r)
	case *s3.DeleteObjectInput:
		deleteObject(r)
	case *s3.DeleteObjectsInput:
		deleteObjects(r)
	case *s3.GetBucketVersioningInput:
		getBucketVersioning(r)
	case *s3.GetObjectInput:
		getObject(r)
	case *s3.ListBucketsInput:
		listBuckets(r)
	case *s3.ListObjectsInput:
		listObjects(r)
	case *s3.PutBucketVersioningInput:
		putBucketVersioning(r)
	case *s3.PutObjectInput:
		putObject(r)
	}
}
