package stakes

import (
	"errors"

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

type stakeBuilder struct {
	met      metadata.MetaData
	fromUser users.User
	toOrg    organizations.Organization
	tok      tokens.Token
	amount   float64
}

func createStakeBuilder() stakes.StakeBuilder {
	out := stakeBuilder{
		met:      nil,
		fromUser: nil,
		toOrg:    nil,
		tok:      nil,
		amount:   0,
	}

	return &out
}

// Create initializes a StakeBuilder instance
func (build *stakeBuilder) Create() stakes.StakeBuilder {
	build.met = nil
	build.fromUser = nil
	build.toOrg = nil
	build.tok = nil
	build.amount = 0
	return build
}

// WithMetaData adds metadata to the StakeBuilder
func (build *stakeBuilder) WithMetaData(met metadata.MetaData) stakes.StakeBuilder {
	build.met = met
	return build
}

// FromUser adds a user to the StakeBuilder
func (build *stakeBuilder) FromUser(usr users.User) stakes.StakeBuilder {
	build.fromUser = usr
	return build
}

// ToOrganization adds an organization to the StakeBuilder
func (build *stakeBuilder) ToOrganization(org organizations.Organization) stakes.StakeBuilder {
	build.toOrg = org
	return build
}

// WithToken adds a token to the StakeBuilder
func (build *stakeBuilder) WithToken(tok tokens.Token) stakes.StakeBuilder {
	build.tok = tok
	return build
}

// WithAmount adds an amount to the StakeBuilder
func (build *stakeBuilder) WithAmount(amount float64) stakes.StakeBuilder {
	build.amount = amount
	return build
}

// Now builds a new Stake instance
func (build *stakeBuilder) Now() (stakes.Stake, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build a stake instance")
	}

	if build.fromUser == nil {
		return nil, errors.New("the fromUse ris mandatory in order to build a stake instance")
	}

	if build.toOrg == nil {
		return nil, errors.New("the toOrganization is mandatory in order to build a stake instance")
	}

	if build.tok == nil {
		return nil, errors.New("the token is mandatory in order to build a stake instance")
	}

	if build.amount == 0 {
		return nil, errors.New("the amount is mandatory in order to build a stake instance")
	}

	out := createStake(build.met.(*concrete_metadata.MetaData), build.fromUser.(*concrete_users.User), build.toOrg.(*concrete_organizations.Organization), build.tok.(*concrete_tokens.Token), build.amount)
	return out, nil
}
