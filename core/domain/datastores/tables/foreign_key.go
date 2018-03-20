package tables

import uuid "github.com/satori/go.uuid"

// ForeignKey represents a table foreign key
type ForeignKey interface {
	GetSchemaName() string
	GetColumnName() string
	GetUUID() *uuid.UUID
}
