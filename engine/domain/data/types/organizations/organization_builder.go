package organizations

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// OrganizationBuilder represents an organization builder
type OrganizationBuilder interface {
	Create() OrganizationBuilder
	WithMetaData(met metadata.MetaData) OrganizationBuilder
	WithUser(usr users.User) OrganizationBuilder
	WithAcceptedToken(tok tokens.Token) OrganizationBuilder
	WithPercentNeededForConcensus(percent float64) OrganizationBuilder
	Now() (Organization, error)
}
