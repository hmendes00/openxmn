package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	uuid "github.com/satori/go.uuid"
)

// MetaDataBuilder represents a chained block metadata builder
type MetaDataBuilder interface {
	Create() MetaDataBuilder
	WithID(id *uuid.UUID) MetaDataBuilder
	WithHashTree(ht hashtrees.HashTree) MetaDataBuilder
	WithPreviousID(prevID *uuid.UUID) MetaDataBuilder
	CreatedOn(crOn time.Time) MetaDataBuilder
	Now() (MetaData, error)
}
