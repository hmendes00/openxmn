package infrastructure

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
)

// Signature represents the concrete user signature
type Signature struct {
	Met *concrete_metadata.MetaData      `json:"metadata"`
	Sig *concrete_cryptography.Signature `json:"signature"`
	Usr *User                            `json:"user"`
}

func createSignature(met *concrete_metadata.MetaData, sig *concrete_cryptography.Signature, usr *User) user.Signature {
	out := Signature{
		Met: met,
		Sig: sig,
		Usr: usr,
	}

	return &out
}

// GetMetaData returns the MetaData
func (sig *Signature) GetMetaData() metadata.MetaData {
	return sig.Met
}

// GetSignature returns the signature
func (sig *Signature) GetSignature() cryptography.Signature {
	return sig.Sig
}

// GetUser returns the user
func (sig *Signature) GetUser() user.User {
	return sig.Usr
}
