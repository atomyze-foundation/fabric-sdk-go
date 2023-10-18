/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package dispatcher

// ConnectionReg is a connection registration
type ConnectionReg struct {
	Eventch chan<- *ConnectionEvent
}
