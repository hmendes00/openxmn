package wealth

import uuid "github.com/satori/go.uuid"

// UpdateToken represents a save token transaction
type UpdateToken struct {
	TokenID *uuid.UUID `json:"token_id"`
	Amount  int        `json:"amount"`
}

// CreateUpdateToken creates a new UpdateToken instance
func CreateUpdateToken(tokenID *uuid.UUID, amount int) *UpdateToken {
	out := UpdateToken{
		TokenID: tokenID,
		Amount:  amount,
	}

	return &out
}

// GetTokenID returns the TokenID
func (tok *UpdateToken) GetTokenID() *uuid.UUID {
	return tok.TokenID
}

// GetAmount returns the amount
func (tok *UpdateToken) GetAmount() int {
	return tok.Amount
}
