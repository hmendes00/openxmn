package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaData represents a chained block metadata
type MetaData interface {
	GetID() *uuid.UUID
	GetIndex() int
	GetPreviousIndex() int
	CreatedOn() time.Time
}
