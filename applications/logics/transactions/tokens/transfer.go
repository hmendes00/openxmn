package tokens

import uuid "github.com/satori/go.uuid"

// Transfer represents a transfer token transaction
type Transfer struct {
	ID       *uuid.UUID `json:"id"`
	TokenID  *uuid.UUID `json:"token_id"`
	FrmUsrID *uuid.UUID `json:"from_user_id"`
	ToOrgID  *uuid.UUID `json:"to_organization_id"`
	Amount   float64    `json:"amount"`
}

// CreateTransfer creates a new Transfer instance
func CreateTransfer(id *uuid.UUID, tokenID *uuid.UUID, fromUserID *uuid.UUID, toOrgID *uuid.UUID, amount float64) *Transfer {
	out := Transfer{
		ID:       id,
		TokenID:  tokenID,
		FrmUsrID: fromUserID,
		ToOrgID:  toOrgID,
		Amount:   amount,
	}

	return &out
}

// GetID returns the ID
func (stk *Transfer) GetID() *uuid.UUID {
	return stk.ID
}

// GetTokenID returns the token ID
func (stk *Transfer) GetTokenID() *uuid.UUID {
	return stk.TokenID
}

// FromUserID returns the from user ID
func (stk *Transfer) FromUserID() *uuid.UUID {
	return stk.FrmUsrID
}

// ToOrganizationID returns the to organization ID
func (stk *Transfer) ToOrganizationID() *uuid.UUID {
	return stk.ToOrgID
}

// GetAmount returns the amount to transfer
func (stk *Transfer) GetAmount() float64 {
	return stk.Amount
}
