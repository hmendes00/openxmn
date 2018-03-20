package tables

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Schema repreents a table schema
type Schema interface {
	GetID() *uuid.UUID
	GetName() string
	GetForeignKeys() []ForeignKey
	LastUpdatedOn() time.Time
	CreatedOn() time.Time
}
