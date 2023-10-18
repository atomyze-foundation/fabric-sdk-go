/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package lbp

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

// LoadBalancePolicy chooses a peer from a set of peers
type LoadBalancePolicy interface {
	Choose(peers []fab.Peer) (fab.Peer, error)
}
