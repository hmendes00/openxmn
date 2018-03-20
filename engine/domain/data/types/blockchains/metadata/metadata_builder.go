package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
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
