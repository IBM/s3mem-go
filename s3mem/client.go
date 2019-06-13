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
	"github.com/aws/aws-sdk-go-v2/private/protocol/restxml"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Client struct {
	s3.Client
}

func New(config aws.Config) *Client {
	svc := &Client{
		Client: *s3.New(config),
	}

	//set handlers
	svc.Handlers = Handlers()
	svc.Handlers.Build.PushBackNamed(restxml.BuildHandler)
	//	svc.Handlers.Unmarshal.PushBackNamed(restxml.UnmarshalHandler)
	// svc.Handlers.UnmarshalMeta.PushBackNamed(restxml.UnmarshalMetaHandler)
	// svc.Handlers.UnmarshalError.PushBackNamed(restxml.UnmarshalErrorHandler)
	svc.AddDebugHandlers()
	return svc
}

// AddDebugHandlers injects debug logging handlers into the service to log request
// debug information.
func (c *Client) AddDebugHandlers() {
	if !c.Config.LogLevel.AtLeast(aws.LogDebug) {
		return
	}

	c.Handlers.Send.PushFrontNamed(aws.NamedHandler{Name: "s3mem.client.LogRequest", Fn: logRequest})
	c.Handlers.Send.PushBackNamed(aws.NamedHandler{Name: "s3mem.client.LogResponse", Fn: logResponse})
}
