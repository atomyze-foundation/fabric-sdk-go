/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: [Default license](LICENSE)
*/

package core

// SigningManager signs object with provided key
type SigningManager interface {
	Sign([]byte, Key) ([]byte, error)
}
