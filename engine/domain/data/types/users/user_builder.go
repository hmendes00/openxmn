package domain

import (
	"time"

	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	uuid "github.com/satori/go.uuid"
)

// UserBuilder represents a User builder
type UserBuilder interface {
	Create() UserBuilder
	WithID(id *uuid.UUID) UserBuilder
	WithMetaData(met metadata.MetaData) UserBuilder
	WithPublicKey(pub cryptography.PublicKey) UserBuilder
	CreatedOn(crOn time.Time) UserBuilder
	LastUpdatedOn(lstUpOn time.Time) UserBuilder
	Now() (User, error)
}
