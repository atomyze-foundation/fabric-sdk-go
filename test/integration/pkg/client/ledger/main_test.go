/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/
package ledger

import (
	"testing"

	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/test/integration"
	"github.com/hyperledger/fabric-sdk-go/test/integration/util/runner"
)

const (
	org1Name      = "Org1"
	org1AdminUser = "Admin"
)

var mainSDK *fabsdk.FabricSDK
var mainTestSetup *integration.BaseSetupImpl
var mainChaincodeID string

func TestMain(m *testing.M) {
	r := runner.NewWithExampleCC()
	r.Initialize()
	mainSDK = r.SDK()
	mainTestSetup = r.TestSetup()
	mainChaincodeID = r.ExampleChaincodeID()

	r.Run(m)
}
