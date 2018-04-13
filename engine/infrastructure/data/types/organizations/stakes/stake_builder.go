package stakes

import (
	"errors"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	stakes "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations/stakes"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_organizations "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/organizations"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

type stakeBuilder struct {
	metBuilderFactory metadata.BuilderFactory
	id                *uuid.UUID
	met               metadata.MetaData
	fromUser          users.User
	toOrg             organizations.Organization
	tok               tokens.Token
	amount            float64
	crOn              *time.Time
	lstUpOn           *time.Time
}

func createStakeBuilder(metBuilderFactory metadata.BuilderFactory) stakes.StakeBuilder {
	out := stakeBuilder{
		metBuilderFactory: metBuilderFactory,
		id:                nil,
		met:               nil,
		fromUser:          nil,
		toOrg:             nil,
		tok:               nil,
		amount:            0,
		crOn:              nil,
		lstUpOn:           nil,
	}

	return &out
}

// Create initializes a StakeBuilder instance
func (build *stakeBuilder) Create() stakes.StakeBuilder {
	build.id = nil
	build.met = nil
	build.fromUser = nil
	build.toOrg = nil
	build.tok = nil
	build.amount = 0
	build.crOn = nil
	build.lstUpOn = nil
	return build
}

// WithID adds an ID to the StakeBuilder
func (build *stakeBuilder) WithID(id *uuid.UUID) stakes.StakeBuilder {
	build.id = id
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

// CreatedOn adds a creation time to the StakeBuilder
func (build *stakeBuilder) CreatedOn(crOn time.Time) stakes.StakeBuilder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds a last updated on time to the StakeBuilder
func (build *stakeBuilder) LastUpdatedOn(lstUpOn time.Time) stakes.StakeBuilder {
	build.lstUpOn = &lstUpOn
	return build
}

// Now builds a new Stake instance
func (build *stakeBuilder) Now() (stakes.Stake, error) {

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

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build a Stake instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build a Stake instance")
		}

		metBuilder := build.metBuilderFactory.Create().Create().WithID(build.id).CreatedOn(*build.crOn)
		if build.lstUpOn != nil {
			metBuilder.LastUpdatedOn(*build.lstUpOn)
		}

		met, metErr := metBuilder.Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	out := createStake(build.met.(*concrete_metadata.MetaData), build.fromUser.(*concrete_users.User), build.toOrg.(*concrete_organizations.Organization), build.tok.(*concrete_tokens.Token), build.amount)
	return out, nil
}
