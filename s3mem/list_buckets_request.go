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

//ListBucketsRequest ...
func (c *Client) ListBucketsRequest(input *s3.ListBucketsInput) s3.ListBucketsRequest {
	if input == nil {
		input = &s3.ListBucketsInput{}
	}
	output := &s3.ListBucketsOutput{}
	req := c.NewRequest(input, output)
	listBuckets := aws.NamedHandler{Name: "S3MemListBuckets", Fn: listBuckets}
	req.Handlers.Send.PushBackNamed(listBuckets)
	return s3.ListBucketsRequest{Request: req, Input: input, Copy: c.ListBucketsRequest}
}

func listBuckets(req *aws.Request) {
	if req.Error != nil {
		return
	}
	req.Data.(*s3.ListBucketsOutput).Buckets = make([]s3.Bucket, 0)
	for _, value := range S3MemBuckets.Buckets {
		req.Data.(*s3.ListBucketsOutput).Buckets = append(req.Data.(*s3.ListBucketsOutput).Buckets, *value.Bucket)
	}
}
