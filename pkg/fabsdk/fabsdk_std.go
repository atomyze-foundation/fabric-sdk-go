//go:build !pprof
// +build !pprof

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

// Package fabsdk enables client usage of a Hyperledger Fabric network.
package fabsdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/metrics"
)

func (sdk *FabricSDK) initMetrics(config *configs) {

	sdk.clientMetrics = &metrics.ClientMetrics{} // empty channel ClientMetrics for standard build.
}
