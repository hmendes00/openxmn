package infrastructure

import (
	commons "github.com/XMNBlockchain/core/packages/controllers/signatures/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

type signature struct {
	sig users.Signature
}

func createSignature() commons.Signature {
	out := signature{
		sig: nil,
	}

	return &out
}

func createSignatureWithUserSignature(sig users.Signature) commons.Signature {
	out := signature{
		sig: sig,
	}

	return &out
}

// HasSignature returns true if there is a user signature, nil otherwise
func (sig *signature) HasSignature() bool {
	return sig.sig != nil
}

// GetSignature returns the user Signature
func (sig *signature) GetSignature() users.Signature {
	return sig.sig
}
