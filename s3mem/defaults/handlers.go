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
package defaults

import (
	"github.com/aws/aws-sdk-go-v2/aws"
)

var checkConfigHandler = aws.NamedHandler{Name: "s3mem.checkConfig", Fn: checkConfig}

func checkConfig(r *aws.Request) {
}
