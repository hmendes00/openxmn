package organizations

import (
	"errors"
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

type organizationBuilder struct {
	metaDataBuilderFactory metadata.BuilderFactory
	id                     *uuid.UUID
	met                    metadata.MetaData
	usr                    users.User
	tok                    tokens.Token
	percent                float64
	crOn                   *time.Time
	lstUpOn                *time.Time
}

func createOrganizationBuilder(metaDataBuilderFactory metadata.BuilderFactory) organizations.OrganizationBuilder {
	out := organizationBuilder{
		metaDataBuilderFactory: metaDataBuilderFactory,
		id:      nil,
		met:     nil,
		usr:     nil,
		tok:     nil,
		percent: 0,
		crOn:    nil,
		lstUpOn: nil,
	}

	return &out
}

// Create initializes the organization builder
func (build *organizationBuilder) Create() organizations.OrganizationBuilder {
	build.id = nil
	build.met = nil
	build.usr = nil
	build.tok = nil
	build.percent = 0
	build.crOn = nil
	build.lstUpOn = nil
	return build
}

// WithID adds an ID to the organization builder
func (build *organizationBuilder) WithID(id *uuid.UUID) organizations.OrganizationBuilder {
	build.id = id
	return build
}

// WithMetaData adds metadata to the organization builder
func (build *organizationBuilder) WithMetaData(met metadata.MetaData) organizations.OrganizationBuilder {
	build.met = met
	return build
}

// WithUser adds a user to the organization builder
func (build *organizationBuilder) WithUser(usr users.User) organizations.OrganizationBuilder {
	build.usr = usr
	return build
}

// WithAcceptedToken adds an accepted token to the organization builder
func (build *organizationBuilder) WithAcceptedToken(tok tokens.Token) organizations.OrganizationBuilder {
	build.tok = tok
	return build
}

// WithPercentNeededForConcensus adds a percent needed for concensus to the organization builder
func (build *organizationBuilder) WithPercentNeededForConcensus(percent float64) organizations.OrganizationBuilder {
	build.percent = percent
	return build
}

// CreatedOn adds a creation time to the organization builder
func (build *organizationBuilder) CreatedOn(crOn time.Time) organizations.OrganizationBuilder {
	build.crOn = &crOn
	return build
}

// LastUpdatedOn adds a last updated on time to the organization builder
func (build *organizationBuilder) LastUpdatedOn(lstUpOn time.Time) organizations.OrganizationBuilder {
	build.lstUpOn = &lstUpOn
	return build
}

// Now builds a new Organization instance
func (build *organizationBuilder) Now() (organizations.Organization, error) {

	if build.usr == nil {
		return nil, errors.New("the user is mandatory in order to build an organization instance")
	}

	if build.tok == nil {
		return nil, errors.New("the token is mandatory in order to build an organization instance")
	}

	if build.percent == 0 {
		return nil, errors.New("the percentNeededForConcensus is mandatory in order to build an organization instance")
	}

	if build.met == nil {
		if build.id == nil {
			return nil, errors.New("the ID is mandatory in order to build an organization instance")
		}

		if build.crOn == nil {
			return nil, errors.New("the creation time is mandatory in order to build an organization instance")
		}

		metBuilder := build.metaDataBuilderFactory.Create().Create().WithID(build.id).CreatedOn(*build.crOn)
		if build.lstUpOn != nil {
			metBuilder.LastUpdatedOn(*build.lstUpOn)
		}

		met, metErr := metBuilder.Now()
		if metErr != nil {
			return nil, metErr
		}

		build.met = met
	}

	out := createOrganization(build.met.(*concrete_metadata.MetaData), build.usr.(*concrete_users.User), build.tok.(*concrete_tokens.Token), build.percent)
	return out, nil
}
