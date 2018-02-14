package domain

import (
	"time"

	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// MetaDataBuilder represents a MetaData builder
type MetaDataBuilder interface {
	Create() MetaDataBuilder
	WithID(id *uuid.UUID) MetaDataBuilder
	WithSignature(sig users.Signature) MetaDataBuilder
	CreatedOn(ts time.Time) MetaDataBuilder
	Now() (MetaData, error)
}
