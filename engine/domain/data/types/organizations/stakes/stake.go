package organizations

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	organizations "github.com/XMNBlockchain/openxmn/engine/domain/data/types/organizations"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Stake represents a stake to an organization
type Stake interface {
	GetMetaData() metadata.MetaData
	FromUser() users.User
	ToOrganization() organizations.Organization
	GetToken() tokens.Token
	GetAmount() float64
}
