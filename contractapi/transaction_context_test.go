// Copyright the Hyperledger Fabric contributors. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package contractapi

import (
	"github.com/wsw365904/newcryptosm/x509"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wsw365904/fabric-chaincode-go/shimtest"
)

// ================================
// Helpers
// ================================

type mockClientIdentity struct{}

func (mci *mockClientIdentity) GetID() (string, error) {
	return "", nil
}

func (mci *mockClientIdentity) GetMSPID() (string, error) {
	return "", nil
}

func (mci *mockClientIdentity) GetAttributeValue(string) (string, bool, error) {
	return "", false, nil
}

func (mci *mockClientIdentity) AssertAttributeValue(string, string) error {
	return nil
}

func (mci *mockClientIdentity) GetX509Certificate() (*x509.Certificate, error) {
	return nil, nil
}

// ================================
// Tests
// ================================

func TestSetStub(t *testing.T) {
	stub := new(shimtest.MockStub)
	stub.TxID = "some ID"

	ctx := TransactionContext{}

	ctx.SetStub(stub)

	assert.Equal(t, stub, ctx.stub, "should have set the same stub as passed")
}

func TestGetStub(t *testing.T) {
	stub := new(shimtest.MockStub)
	stub.TxID = "some ID"

	ctx := TransactionContext{}
	ctx.stub = stub

	assert.Equal(t, stub, ctx.GetStub(), "should have returned same stub as set")
}

func TestSetClientIdentity(t *testing.T) {
	ci := new(mockClientIdentity)

	ctx := TransactionContext{}

	ctx.SetClientIdentity(ci)

	assert.Equal(t, ci, ctx.clientIdentity, "should have set the same client identity as passed")
}

func TestGetClientIdentity(t *testing.T) {
	ci := new(mockClientIdentity)

	ctx := TransactionContext{}
	ctx.clientIdentity = ci

	assert.Equal(t, ci, ctx.GetClientIdentity(), "should have returned same client identity as set")
}