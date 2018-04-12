package wealth

import uuid "github.com/satori/go.uuid"

// UpdateToken represents a save token transaction
type UpdateToken struct {
	TokenID *uuid.UUID `json:"token_id"`
	Symbol  string     `json:"symbol"`
}

// CreateUpdateToken creates a new UpdateToken instance
func CreateUpdateToken(tokenID *uuid.UUID, symbol string) *UpdateToken {
	out := UpdateToken{
		TokenID: tokenID,
		Symbol:  symbol,
	}

	return &out
}

// GetTokenID returns the TokenID
func (tok *UpdateToken) GetTokenID() *uuid.UUID {
	return tok.TokenID
}

// GetSymbol returns the symbol
func (tok *UpdateToken) GetSymbol() string {
	return tok.Symbol
}
