package infrastructure

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	user "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// Signatures represents a concrete Signatures implementation
type Signatures struct {
	Met  *concrete_metadata.MetaData `json:"metadata"`
	Sigs []*Signature                `json:"signatures"`
}

func createSignatures(met *concrete_metadata.MetaData, sigs []*Signature) user.Signatures {
	out := Signatures{
		Met:  met,
		Sigs: sigs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (sigs *Signatures) GetMetaData() metadata.MetaData {
	return sigs.Met
}

// GetSignatures returns the []Signature
func (sigs *Signatures) GetSignatures() []user.Signature {
	out := []user.Signature{}
	for _, oneSig := range sigs.Sigs {
		out = append(out, oneSig)
	}

	return out
}
