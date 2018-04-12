package safes

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
)

// Safe represents a concrete safe implementation
type Safe struct {
	Met  *concrete_metadata.MetaData   `json:"metadata"`
	Ciph *concrete_cryptography.Cipher `json:"cipher"`
}

func createSafe(met *concrete_metadata.MetaData, ciph *concrete_cryptography.Cipher) safes.Safe {
	out := Safe{
		Met:  met,
		Ciph: ciph,
	}

	return &out
}

// GetMetaData returns the metadata
func (saf *Safe) GetMetaData() metadata.MetaData {
	return saf.Met
}

// GetCipher returns the cipher
func (saf *Safe) GetCipher() cryptography.Cipher {
	return saf.Ciph
}
