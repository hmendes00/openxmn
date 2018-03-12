package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	uuid "github.com/satori/go.uuid"
)

// MetaDataBuilder represents a MetaData builder
type MetaDataBuilder interface {
	Create() MetaDataBuilder
	WithID(id *uuid.UUID) MetaDataBuilder
	WithHashTree(ht hashtrees.HashTree) MetaDataBuilder
	CreatedOn(ts time.Time) MetaDataBuilder
	Now() (MetaData, error)
}
