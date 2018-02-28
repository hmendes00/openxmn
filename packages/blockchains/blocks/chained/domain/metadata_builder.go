package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaDataBuilder represents a chained block metadata builder
type MetaDataBuilder interface {
	Create() MetaDataBuilder
	WithID(id *uuid.UUID) MetaDataBuilder
	WithIndex(index int) MetaDataBuilder
	WithPreviousIndex(prevIndex int) MetaDataBuilder
	CreatedOn() time.Time
	Now() (MetaData, error)
}
