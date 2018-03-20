package tables

import uuid "github.com/satori/go.uuid"

// SchemaRepository represents a schema repository
type SchemaRepository interface {
	RetrieveByID(id *uuid.UUID) (Schema, error)
}
