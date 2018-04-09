package metadata

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Builder represents a metadata builder
type Builder interface {
	Create() Builder
	WithID(id *uuid.UUID) Builder
	CreatedOn(crOn time.Time) Builder
	LastUpdatedOn(lstUpOn time.Time) Builder
	Now() (MetaData, error)
}
