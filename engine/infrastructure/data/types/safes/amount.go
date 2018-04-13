package safes

import (
	safes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/safes"
	uuid "github.com/satori/go.uuid"
)

// Amount represents a concrete amount representation
type Amount struct {
	TokID  *uuid.UUID `json:"token_id"`
	Amount float64    `json:"amount"`
}

func createAmount(tokID *uuid.UUID, amount float64) safes.Amount {
	out := Amount{
		TokID:  tokID,
		Amount: amount,
	}

	return &out
}

// GetTokenID returns the tokenID
func (am *Amount) GetTokenID() *uuid.UUID {
	return am.TokID
}

// GetAmount returns the amount
func (am *Amount) GetAmount() float64 {
	return am.Amount
}
