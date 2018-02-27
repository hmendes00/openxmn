package tokens

import uuid "github.com/satori/go.uuid"

// Stake represents a stake token transaction
type Stake struct {
	ID       *uuid.UUID `json:"id"`
	TokenID  *uuid.UUID `json:"token_id"`
	FrmUsrID *uuid.UUID `json:"from_user_id"`
	ToOrgID  *uuid.UUID `json:"to_organization_id"`
	Amount   float64    `json:"amount"`
}

// CreateStake creates a new Stake instance
func CreateStake(id *uuid.UUID, tokenID *uuid.UUID, fromUserID *uuid.UUID, toOrgID *uuid.UUID, amount float64) *Stake {
	out := Stake{
		ID:       id,
		TokenID:  tokenID,
		FrmUsrID: fromUserID,
		ToOrgID:  toOrgID,
		Amount:   amount,
	}

	return &out
}

// GetID returns the ID
func (stk *Stake) GetID() *uuid.UUID {
	return stk.ID
}

// GetTokenID returns the token ID
func (stk *Stake) GetTokenID() *uuid.UUID {
	return stk.TokenID
}

// FromUserID returns the from user ID
func (stk *Stake) FromUserID() *uuid.UUID {
	return stk.FrmUsrID
}

// ToOrganizationID returns the to organization ID
func (stk *Stake) ToOrganizationID() *uuid.UUID {
	return stk.ToOrgID
}

// GetAmount returns the amount to stake
func (stk *Stake) GetAmount() float64 {
	return stk.Amount
}
