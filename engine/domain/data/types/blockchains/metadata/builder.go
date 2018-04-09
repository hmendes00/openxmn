package metadata

import (
	"time"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	uuid "github.com/satori/go.uuid"
)

// Builder represents a MetaData builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	WithHashTree(ht hashtrees.HashTree) Builder
	CreatedOn(ts time.Time) Builder
	Now() (MetaData, error)
}
