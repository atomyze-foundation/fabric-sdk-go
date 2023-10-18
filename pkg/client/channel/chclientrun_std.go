// +build !pprof

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package channel

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel/invoke"
)

func callQuery(cc *Client, request Request, options ...RequestOption) (Response, error) {
	return cc.InvokeHandler(invoke.NewQueryHandler(), request, options...)
}

func callExecute(cc *Client, request Request, options ...RequestOption) (Response, error) {
	return cc.InvokeHandler(invoke.NewExecuteHandler(), request, options...)
}
