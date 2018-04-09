package organizations

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	tokens "github.com/XMNBlockchain/openxmn/engine/domain/data/types/tokens"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Organization represents an organization
type Organization interface {
	GetMetaData() metadata.MetaData
	GetUser() users.User
	GetAcceptedToken() tokens.Token
	GetPercentNeededForConcensus() float64
}
