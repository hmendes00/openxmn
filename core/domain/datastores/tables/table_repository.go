package tables

import uuid "github.com/satori/go.uuid"

// TableRepository represents a table repository
type TableRepository interface {
	RetrieveByID(dirPath string, id *uuid.UUID) (Table, error)
}
