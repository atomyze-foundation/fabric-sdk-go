/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package fabsdk

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/core/logging/api"
	sdkApi "github.com/hyperledger/fabric-sdk-go/pkg/fabsdk/api"
)

// pkgSuite provides the package factories that create clients and providers
type pkgSuite interface {
	Core() (sdkApi.CoreProviderFactory, error)
	MSP() (sdkApi.MSPProviderFactory, error)
	Service() (sdkApi.ServiceProviderFactory, error)
	Logger() (api.LoggerProvider, error)
}
