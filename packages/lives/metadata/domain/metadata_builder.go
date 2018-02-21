package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaDataBuilder represents a MetaData builder
type MetaDataBuilder interface {
	Create() MetaDataBuilder
	WithID(id *uuid.UUID) MetaDataBuilder
	CreatedOn(ts time.Time) MetaDataBuilder
	Now() (MetaData, error)
}
