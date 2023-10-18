/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package channel

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

const (
	cscc            = "cscc"
	csccConfigBlock = "GetConfigBlock"
)

func createConfigBlockInvokeRequest(channelID string) fab.ChaincodeInvokeRequest {
	cir := fab.ChaincodeInvokeRequest{
		ChaincodeID: cscc,
		Fcn:         csccConfigBlock,
		Args:        [][]byte{[]byte(channelID)},
	}
	return cir
}
