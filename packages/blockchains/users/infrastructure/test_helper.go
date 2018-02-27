package infrastructure

import (
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	uuid "github.com/satori/go.uuid"
)

// CreateUserForTests creates a User for tests
func CreateUserForTests(t *testing.T) *User {
	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)

	user := createUser(&id, pk)
	return user.(*User)
}

// CreateSignatureForTests creates a Signature for tests
func CreateSignatureForTests(t *testing.T) *Signature {
	//variables:
	sig := concrete_cryptography.CreateSignatureForTests(t)
	usr := CreateUserForTests(t)

	userSig := createSignature(sig, usr)
	return userSig.(*Signature)
}
