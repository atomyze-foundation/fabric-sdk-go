//go:build testing
// +build testing

/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package logging

import "sync"

// UnsafeReset allows reinitialization of the logger provider.
// This method is intended to enable tests and should not be called.
func UnsafeReset() {
	loggerProviderInstance = nil
	loggerProviderOnce = sync.Once{}
}
