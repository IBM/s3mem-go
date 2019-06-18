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

package example

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/s3iface"
)

func GetObject(client s3iface.ClientAPI, bucket *string, key *string) ([]byte, error) {
	//Create a request
	req := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	//Send the request
	getObjectOutput, err := req.Send(context.TODO())
	if err != nil {
		return nil, err
	}
	//Read body
	buf := new(bytes.Buffer)
	buf.ReadFrom(getObjectOutput.Body)
	newBytes := buf.Bytes()
	return newBytes, nil
}

func GetBuckets(client s3iface.ClientAPI) ([]string, error) {
	//Create the request
	req := client.ListBucketsRequest(&s3.ListBucketsInput{})
	//Send the request
	listBucketsOutput, err := req.Send(context.Background())
	if err != nil {
		return nil, err
	}
	buckets := make([]string, len(listBucketsOutput.Buckets))
	for i, v := range listBucketsOutput.Buckets {
		buckets[i] = *v.Name
	}
	return buckets, nil
}
