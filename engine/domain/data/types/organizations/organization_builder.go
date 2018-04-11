package organizations

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// OrganizationBuilder represents an organization builder
type OrganizationBuilder interface {
	Create() OrganizationBuilder
	WithID(id *uuid.UUID) OrganizationBuilder
	WithMetaData(met metadata.MetaData) OrganizationBuilder
	WithUser(usr users.User) OrganizationBuilder
	WithAcceptedToken(tok tokens.Token) OrganizationBuilder
	WithPercentNeededForConcensus(percent float64) OrganizationBuilder
	CreatedOn(crOn time.Time) OrganizationBuilder
	LastUpdatedOn(lstUpOn time.Time) OrganizationBuilder
	Now() (Organization, error)
}
