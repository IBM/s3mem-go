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
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func (c *Client) NewRequest(operation *aws.Operation, params interface{}, data interface{}) *aws.Request {

	// TODO improve this experiance for config copy?
	cfg := c.Config.Copy()
	method := operation.HTTPMethod
	if method == "" {
		method = "POST"
	}

	httpReq, _ := http.NewRequest(method, "", nil)

	metadata := c.Metadata
	// if metadata.EndpointsID == "" {
	// 	metadata.EndpointsID = s3.EndpointsID
	// }

	// if cfg.Region == nil {
	// 	cfg.Region = endpoints.UsEast1RegionID
	// }

	// if cfg.EndpointResolver == nil {
	// 	cfg.EndpointResolver = NewDefaultResolver()
	// }

	// TODO need better way of handling this error... NewRequest should return error.
	endpoint, err := cfg.EndpointResolver.ResolveEndpoint(metadata.EndpointsID, cfg.Region)
	if err == nil {
		// TODO so ugly
		metadata.Endpoint = endpoint.URL
		if len(endpoint.SigningName) > 0 && !endpoint.SigningNameDerived {
			metadata.SigningName = endpoint.SigningName
		}
		if len(endpoint.SigningRegion) > 0 {
			metadata.SigningRegion = endpoint.SigningRegion
		}

		httpReq.URL, err = url.Parse(endpoint.URL + operation.HTTPPath)
		if err != nil {
			httpReq.URL = &url.URL{}
			err = awserr.New("InvalidEndpointURL", "invalid endpoint uri", err)
		}
	}

	handlers := c.Handlers

	r := &aws.Request{
		Config: aws.Config{
			Region:      cfg.Region,
			Credentials: cfg.Credentials,
			Handlers:    cfg.Handlers,
			LogLevel:    cfg.LogLevel,
			Logger:      cfg.Logger,
		},
		Metadata: metadata,
		Handlers: handlers.Copy(),

		Time:         time.Now(),
		ExpireTime:   0,
		Operation:    operation,
		HTTPRequest:  httpReq,
		HTTPResponse: &http.Response{},
		Body:         nil,
		Params:       params,
		Error:        nil,
		Data:         data,
	}
	r.SetBufferBody([]byte{})

	return r
}

func NewDefaultResolver() aws.EndpointResolver {
	defaultResolver := endpoints.NewDefaultResolver()
	myCustomResolver := func(service, region string) (aws.Endpoint, error) {
		if service == s3.EndpointsID {
			return aws.Endpoint{
				URL: S3MemEndpointsID,
			}, nil
		}
		return defaultResolver.ResolveEndpoint(service, region)
	}
	endpointResolver := aws.EndpointResolverFunc(myCustomResolver)
	return endpointResolver
}
