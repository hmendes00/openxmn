package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/stores/files"
)

// Signature represents a concrete stored user signature implementation
type Signature struct {
	Met *concrete_stored_files.File `json:"metadata"`
	Sig *concrete_stored_files.File `json:"signature"`
	Usr *User                       `json:"user"`
}

func createSignature(met *concrete_stored_files.File, sig *concrete_stored_files.File, usr *User) stored_users.Signature {
	out := Signature{
		Met: met,
		Sig: sig,
		Usr: usr,
	}

	return &out
}

// GetMetaData returns the MetaData
func (sig *Signature) GetMetaData() stored_files.File {
	return sig.Met
}

// GetSignature returns the signature
func (sig *Signature) GetSignature() stored_files.File {
	return sig.Sig
}

// GetUser returns the User
func (sig *Signature) GetUser() stored_users.User {
	return sig.Usr
}
