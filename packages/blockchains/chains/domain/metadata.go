package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaData represents the chain metadata
type MetaData interface {
	GetID() *uuid.UUID
	GetHeight() int
	CreatedOn() time.Time
	LastUpdatedOn() time.Time
}
