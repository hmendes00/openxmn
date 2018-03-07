package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
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
