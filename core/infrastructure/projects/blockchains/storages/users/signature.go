package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

type signature struct {
	met stored_files.File
	sig stored_files.File
	usr stored_users.User
}

func createSignature(met stored_files.File, sig stored_files.File, usr stored_users.User) stored_users.Signature {
	out := signature{
		met: met,
		sig: sig,
		usr: usr,
	}

	return &out
}

// GetMetaData returns the MetaData
func (sig *signature) GetMetaData() stored_files.File {
	return sig.met
}

// GetSignature returns the signature
func (sig *signature) GetSignature() stored_files.File {
	return sig.sig
}

// GetUser returns the User
func (sig *signature) GetUser() stored_users.User {
	return sig.usr
}
