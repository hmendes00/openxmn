package organizations

import (
	"errors"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	concrete_tokens "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/tokens"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
)

type organizationBuilder struct {
	met     metadata.MetaData
	usr     users.User
	tok     tokens.Token
	percent float64
}

func createOrganizationBuilder() organizations.OrganizationBuilder {
	out := organizationBuilder{
		met:     nil,
		usr:     nil,
		tok:     nil,
		percent: 0,
	}

	return &out
}

// Create initializes the organization builder
func (build *organizationBuilder) Create() organizations.OrganizationBuilder {
	build.met = nil
	build.usr = nil
	build.tok = nil
	build.percent = 0
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

// Now builds a new Organization instance
func (build *organizationBuilder) Now() (organizations.Organization, error) {
	if build.met == nil {
		return nil, errors.New("the metadata is mandatory in order to build an organization instance")
	}

	if build.usr == nil {
		return nil, errors.New("the user is mandatory in order to build an organization instance")
	}

	if build.tok == nil {
		return nil, errors.New("the token is mandatory in order to build an organization instance")
	}

	if build.percent == 0 {
		return nil, errors.New("the percentNeededForConcensus is mandatory in order to build an organization instance")
	}

	out := createOrganization(build.met.(*concrete_metadata.MetaData), build.usr.(*concrete_users.User), build.tok.(*concrete_tokens.Token), build.percent)
	return out, nil
}
