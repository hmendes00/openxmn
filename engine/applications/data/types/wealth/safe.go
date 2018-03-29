package wealth

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
	cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
)

// Safe represents a safe containing a specific amount of token
type Safe struct {
	Met    *metadata.MetaData   `json:"metadata"`
	Tok    *Token               `json:"token"`
	Cipher *cryptography.Cipher `json:"cipher"`
}

// CreateSafe creates a new Safe instance
func CreateSafe(met *metadata.MetaData, tok *Token, ci *cryptography.Cipher) *Safe {
	out := Safe{
		Met:    met,
		Tok:    tok,
		Cipher: ci,
	}

	return &out
}

// GetMetaData returns the metadata
func (saf *Safe) GetMetaData() *metadata.MetaData {
	return saf.Met
}

// GetToken returns the token
func (saf *Safe) GetToken() *Token {
	return saf.Tok
}

// GetCipher returns the cipher
func (saf *Safe) GetCipher() *cryptography.Cipher {
	return saf.Cipher
}
