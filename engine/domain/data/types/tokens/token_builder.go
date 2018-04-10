package tokens

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// TokenBuilder represents a token builder
type TokenBuilder interface {
	Create() TokenBuilder
	WithID(id *uuid.UUID) TokenBuilder
	WithMetaData(met metadata.MetaData) TokenBuilder
	WithCreator(creator users.User) TokenBuilder
	WithSymbol(symbol string) TokenBuilder
	WithAmount(amount int) TokenBuilder
	CreatedOn(crOn time.Time) TokenBuilder
	Now() (Token, error)
}
