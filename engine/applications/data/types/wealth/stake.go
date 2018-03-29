package wealth

import (
	"github.com/XMNBlockchain/openxmn/engine/applications/data/types/metadata"
)

// Stake represents tokens put at stake to an organization, by an entity
type Stake struct {
	Met       *metadata.MetaData `json:"metadata"`
	ToOrg     *Organization      `json:"to_organization"`
	FrmEntity *Entity            `json:"from_entity"`
	Tok       *Token             `json:"token"`
	Amount    float64            `json:"amount"`
}

// CreateStake creates a new stake instance
func CreateStake(met *metadata.MetaData, toOrg *Organization, frmEntity *Entity, tok *Token, amount float64) *Stake {
	out := Stake{
		Met:       met,
		ToOrg:     toOrg,
		FrmEntity: frmEntity,
		Tok:       tok,
		Amount:    amount,
	}

	return &out
}

// GetMetaData returns the metadata
func (stk *Stake) GetMetaData() *metadata.MetaData {
	return stk.Met
}

// ToOrganization returns the organization to which the stake is created on
func (stk *Stake) ToOrganization() *Organization {
	return stk.ToOrg
}

// FromEntiity returns the entity that staked token to an organization
func (stk *Stake) FromEntiity() *Entity {
	return stk.FrmEntity
}

// GetToken returns the token
func (stk *Stake) GetToken() *Token {
	return stk.Tok
}

// GetAmount returns the amount
func (stk *Stake) GetAmount() float64 {
	return stk.Amount
}
