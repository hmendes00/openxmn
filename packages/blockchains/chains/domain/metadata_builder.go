package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaDataBuilder represents the metadata builder
type MetaDataBuilder interface {
	Create() MetaDataBuilder
	WithID(id *uuid.UUID) MetaDataBuilder
	WithHeight(height int) MetaDataBuilder
	CreatedOn(crOn time.Time) MetaDataBuilder
	LastUpdatedOn(lastOn time.Time) MetaDataBuilder
	Now() (MetaData, error)
}
