package wealth

import uuid "github.com/satori/go.uuid"

// StakeToken represents a stake token transaction
type StakeToken struct {
	StakeID *uuid.UUID `json:"stake_id"`
	TokenID *uuid.UUID `json:"token_id"`
	OrgID   *uuid.UUID `json:"organization_id"`
	Amount  float64    `json:"amount"`
}

// CreateStakeToken create a new StakeToken instance
func CreateStakeToken(stakeID *uuid.UUID, tokenID *uuid.UUID, organizationID *uuid.UUID, amount float64) *StakeToken {
	out := StakeToken{
		StakeID: stakeID,
		TokenID: tokenID,
		OrgID:   organizationID,
		Amount:  amount,
	}

	return &out
}

// GetStakeID returns the StakeID
func (tok *StakeToken) GetStakeID() *uuid.UUID {
	return tok.StakeID
}

// GetTokenID returns the TokenID
func (tok *StakeToken) GetTokenID() *uuid.UUID {
	return tok.TokenID
}

// GetOrganizationID returns the OrganizationID
func (tok *StakeToken) GetOrganizationID() *uuid.UUID {
	return tok.OrgID
}

// GetAmount returns the amount
func (tok *StakeToken) GetAmount() float64 {
	return tok.Amount
}
