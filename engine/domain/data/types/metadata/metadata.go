package metadata

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaData represents the metadata of a type
type MetaData interface {
	GetID() *uuid.UUID
	CreatedOn() time.Time
	LastUpdatedOn() time.Time
}
