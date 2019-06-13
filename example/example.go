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
