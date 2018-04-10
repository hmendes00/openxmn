package wealth

import uuid "github.com/satori/go.uuid"

// InsertToken represents a save token transaction
type InsertToken struct {
	TokenID *uuid.UUID `json:"token_id"`
	Symbol  string     `json:"symbol"`
	Amount  int        `json:"amount"`
}

// CreateInsertToken creates a new InsertToken instance
func CreateInsertToken(tokenID *uuid.UUID, symbol string, amount int) *InsertToken {
	out := InsertToken{
		TokenID: tokenID,
		Symbol:  symbol,
		Amount:  amount,
	}

	return &out
}

// GetTokenID returns the TokenID
func (tok *InsertToken) GetTokenID() *uuid.UUID {
	return tok.TokenID
}

// GetSymbol returns the symbol
func (tok *InsertToken) GetSymbol() string {
	return tok.Symbol
}

// GetAmount returns the amount
func (tok *InsertToken) GetAmount() int {
	return tok.Amount
}
