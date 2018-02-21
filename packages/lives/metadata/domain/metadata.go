package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// MetaData represents an object metadata
type MetaData interface {
	GetID() *uuid.UUID
	CreatedOn() time.Time
}
