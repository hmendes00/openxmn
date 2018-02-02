package infrastructure

import (
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	user "github.com/XMNBlockchain/core/packages/users/domain"
)

// Signature represents the concrete user signature
type Signature struct {
	Sig *concrete_cryptography.Signature `json:"signature"`
	Usr *User                            `json:"user"`
}

func createSignature(sig *concrete_cryptography.Signature, usr *User) user.Signature {
	out := Signature{
		Sig: sig,
		Usr: usr,
	}

	return &out
}

// GetSig returns the signature
func (sig *Signature) GetSig() cryptography.Signature {
	return sig.Sig
}

// GetUser returns the user
func (sig *Signature) GetUser() user.User {
	return sig.Usr
}
