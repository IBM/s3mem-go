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
	"net/http"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
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

	if cfg.EndpointResolver == nil {
		cfg.EndpointResolver = endpoints.NewDefaultResolver()
	}

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
