package tokens

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithMetaData(met metadata.MetaData) TokenBuilder
	WithCreator(creator users.User) TokenBuilder
	WithSymbol(symbol string) TokenBuilder
	WithAmount(amount int) TokenBuilder
	Now() (Token, error)
}
