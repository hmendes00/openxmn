package wealth

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	uuid "github.com/satori/go.uuid"
)

// InsertToken represents a save token transaction
type InsertToken struct {
	TokenID   *uuid.UUID                    `json:"token_id"`
	CreatorID *uuid.UUID                    `json:"creator_id"`
	Symbol    string                        `json:"symbol"`
	SafeID    *uuid.UUID                    `json:"safe_id"`
	Cipher    *concrete_cryptography.Cipher `json:"cipher"`
}

// CreateInsertToken creates a new InsertToken instance
func CreateInsertToken(tokenID *uuid.UUID, creatorID *uuid.UUID, symbol string, safeID *uuid.UUID, cipher *concrete_cryptography.Cipher) *InsertToken {
	out := InsertToken{
		TokenID:   tokenID,
		CreatorID: creatorID,
		Symbol:    symbol,
		SafeID:    safeID,
		Cipher:    cipher,
	}

	return &out
}

// GetTokenID returns the TokenID
func (tok *InsertToken) GetTokenID() *uuid.UUID {
	return tok.TokenID
}

// GetCreatorID returns the creatorID
func (tok *InsertToken) GetCreatorID() *uuid.UUID {
	return tok.CreatorID
}

// GetSymbol returns the symbol
func (tok *InsertToken) GetSymbol() string {
	return tok.Symbol
}

// GetSafeID returns the safeID
func (tok *InsertToken) GetSafeID() *uuid.UUID {
	return tok.SafeID
}

// GetCipher returns the cipher
func (tok *InsertToken) GetCipher() cryptography.Cipher {
	return tok.Cipher
}
