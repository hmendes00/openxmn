package safes

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
)

// Safe represents a concrete safe implementation
type Safe struct {
	Met  *concrete_metadata.MetaData   `json:"metadata"`
	Tok  *concrete_tokens.Token        `json:"token"`
	Ciph *concrete_cryptography.Cipher `json:"cipher"`
}

func createSafe(met *concrete_metadata.MetaData, tok *concrete_tokens.Token, ciph *concrete_cryptography.Cipher) safes.Safe {
	out := Safe{
		Met:  met,
		Tok:  tok,
		Ciph: ciph,
	}

	return &out
}

// GetMetaData returns the metadata
func (saf *Safe) GetMetaData() metadata.MetaData {
	return saf.Met
}

// GetToken returns the token
func (saf *Safe) GetToken() tokens.Token {
	return saf.Tok
}

// GetCipher returns the cipher
func (saf *Safe) GetCipher() cryptography.Cipher {
	return saf.Ciph
}
