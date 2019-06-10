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
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
)

func (c *Client) NewRequest(params interface{}, data interface{}) *aws.Request {

	// TODO improve this experiance for config copy?
	cfg := c.Config.Copy()
	metadata := c.Metadata

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

		Time:        time.Now(),
		ExpireTime:  0,
		Operation:   &aws.Operation{},
		HTTPRequest: &http.Request{},
		Body:        nil,
		Params:      params,
		Error:       nil,
		Data:        data,
	}

	return r
}
