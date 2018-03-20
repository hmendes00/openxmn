package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
)

// Signatures represents a concrete stored signatures implementation
type Signatures struct {
	Met  *concrete_stored_files.File `json:"metadata"`
	Sigs []*Signature                `json:"signatures"`
}

func createSignatures(met *concrete_stored_files.File, sigs []*Signature) stored_users.Signatures {
	out := Signatures{
		Met:  met,
		Sigs: sigs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (sigs *Signatures) GetMetaData() stored_files.File {
	return sigs.Met
}

// GetSignatures returns the []Signature
func (sigs *Signatures) GetSignatures() []stored_users.Signature {
	out := []stored_users.Signature{}
	for _, oneSig := range sigs.Sigs {
		out = append(out, oneSig)
	}
	return out
}
