package stakes

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// StakeBuilder represents a stake builder
type StakeBuilder interface {
	Create() StakeBuilder
	WithID(id *uuid.UUID) StakeBuilder
	WithMetaData(met metadata.MetaData) StakeBuilder
	FromUser(usr users.User) StakeBuilder
	ToOrganization(org organizations.Organization) StakeBuilder
	WithToken(tok tokens.Token) StakeBuilder
	WithAmount(amount float64) StakeBuilder
	CreatedOn(crOn time.Time) StakeBuilder
	LastUpdatedOn(lstUpOn time.Time) StakeBuilder
	Now() (Stake, error)
}
