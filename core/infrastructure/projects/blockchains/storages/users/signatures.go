package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

type signatures struct {
	met  stored_files.File
	sigs []stored_users.Signature
}

func createSignatures(met stored_files.File, sigs []stored_users.Signature) stored_users.Signatures {
	out := signatures{
		met:  met,
		sigs: sigs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (sigs *signatures) GetMetaData() stored_files.File {
	return sigs.met
}

// GetSignatures returns the []Signature
func (sigs *signatures) GetSignatures() []stored_users.Signature {
	return sigs.sigs
}
