package tokens

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Token represents a token
type Token interface {
	GetMetaData() metadata.MetaData
	GetCreator() users.User
	GetSymbol() string
	GetAmount() int
}
