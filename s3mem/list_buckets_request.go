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
	"encoding/json"
	"io/ioutil"
	"sort"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

const opListBuckets = "ListBuckets"

//ListBucketsRequest ...
func (c *Client) ListBucketsRequest(input *s3.ListBucketsInput) s3.ListBucketsRequest {
	if input == nil {
		input = &s3.ListBucketsInput{}
	}
	output := &s3.ListBucketsOutput{}
	op := &aws.Operation{
		Name:       opListBuckets,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}
	req := c.NewRequest(op, input, output)
	return s3.ListBucketsRequest{Request: req, Input: input, Copy: c.ListBucketsRequest}
}

func listBuckets(req *aws.Request) {
	if req.Error != nil {
		return
	}
	req.Data.(*s3.ListBucketsOutput).Buckets = make([]s3.Bucket, 0)
	var keys []string
	for k := range S3MemBuckets.Buckets {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		req.Data.(*s3.ListBucketsOutput).Buckets = append(req.Data.(*s3.ListBucketsOutput).Buckets, *S3MemBuckets.Buckets[k].Bucket)
	}
	//This is needed just to logResponse when requested
	body, _ := json.MarshalIndent(req.Data.(*s3.ListBucketsOutput), "", "  ")
	req.HTTPResponse.Body = ioutil.NopCloser(bytes.NewReader(body))
}
