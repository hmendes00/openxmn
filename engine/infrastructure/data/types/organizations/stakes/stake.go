package stakes

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_organizations "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

// Stake represents a concrete stake implementation
type Stake struct {
	Met     *concrete_metadata.MetaData          `json:"metadata"`
	FrmUser *concrete_users.User                 `json:"from_user"`
	ToOrg   *concrete_organizations.Organization `json:"to_organization"`
	Tok     *concrete_tokens.Token               `json:"token"`
	Amount  float64                              `json:"amount"`
}

func createStake(met *concrete_metadata.MetaData, fromUser *concrete_users.User, toOrganization *concrete_organizations.Organization, token *concrete_tokens.Token, amount float64) stakes.Stake {
	out := Stake{
		Met:     met,
		FrmUser: fromUser,
		ToOrg:   toOrganization,
		Tok:     token,
		Amount:  amount,
	}

	return &out
}

// GetMetaData returns the metadata
func (stk *Stake) GetMetaData() metadata.MetaData {
	return stk.Met
}

// FromUser returns the user
func (stk *Stake) FromUser() users.User {
	return stk.FrmUser
}

// ToOrganization returns the organization
func (stk *Stake) ToOrganization() organizations.Organization {
	return stk.ToOrg
}

// GetToken returns the token
func (stk *Stake) GetToken() tokens.Token {
	return stk.Tok
}

// GetAmount returns the amount
func (stk *Stake) GetAmount() float64 {
	return stk.Amount
}
