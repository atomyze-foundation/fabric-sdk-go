// +build testing

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package dynamicdiscovery

import (
	contextAPI "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/fab/discovery"
)

// SetClientProvider overrides the discovery client provider for unit tests
func SetClientProvider(provider func(ctx contextAPI.Client) (discovery.Client, error)) {
	clientProvider = provider
}
