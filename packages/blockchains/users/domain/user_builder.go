package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
	uuid "github.com/satori/go.uuid"
)

// UserBuilder represents a User builder
type UserBuilder interface {
	Create() UserBuilder
	WithID(id uuid.UUID) UserBuilder
	WithMetaData(met metadata.MetaData) UserBuilder
	WithPublicKey(pub cryptography.PublicKey) UserBuilder
	CreatedOn(crOn time.Time) UserBuilder
	Now() (User, error)
}
