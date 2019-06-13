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
)

const S3MemEndpointsID = "s3mem"

func Handlers() aws.Handlers {
	var handlers aws.Handlers

	// handlers.Validate.PushBackNamed(defaults.ValidateEndpointHandler)
	// handlers.Validate.PushBackNamed(defaults.ValidateParametersHandler)
	// handlers.Validate.AfterEachFn = aws.HandlerListStopOnError
	// handlers.Build.PushBackNamed(defaults.SDKVersionUserAgentHandler)
	// handlers.Build.PushBackNamed(defaults.AddHostExecEnvUserAgentHander)
	// handlers.Build.AfterEachFn = aws.HandlerListStopOnError
	// handlers.Sign.PushBackNamed(defaults.BuildContentLengthHandler)
	// handlers.Send.PushBackNamed(defaults.ValidateReqSigHandler)
	// handlers.Send.PushBackNamed(SendHandler)
	handlers.Send.PushBackNamed(sendHandler)
	// handlers.AfterRetry.PushBackNamed(defaults.AfterRetryHandler)
	// handlers.ValidateResponse.PushBackNamed(defaults.ValidateResponseHandler)

	return handlers
}
