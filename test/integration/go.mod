// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/hyperledger/fabric-sdk-go/test/integration

replace github.com/hyperledger/fabric-sdk-go => ../../

require (
	github.com/golang/protobuf v1.5.3
	github.com/hyperledger/fabric-config v0.1.0
	github.com/hyperledger/fabric-protos-go v0.3.1
	github.com/hyperledger/fabric-sdk-go v0.0.0-00010101000000-000000000000
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.3
	google.golang.org/grpc v1.59.0
	google.golang.org/protobuf v1.31.0 // indirect
)

go 1.14
