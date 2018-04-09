package organizations

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// StakeBuilder represents a stake builder
type StakeBuilder interface {
	Create() StakeBuilder
	WithMetaData(met metadata.MetaData) StakeBuilder
	FromUser(usr users.User) StakeBuilder
	ToOrganization(org organizations.Organization) StakeBuilder
	WithToken(tok tokens.Token) StakeBuilder
	WithAmount(amount float64) StakeBuilder
	Now() (Stake, error)
}
