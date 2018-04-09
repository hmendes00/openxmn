package files

import (
	"hash"

	uuid "github.com/satori/go.uuid"
)

// FileBuilder represents a file builder
type FileBuilder interface {
	Create() FileBuilder
	WithID(id *uuid.UUID) FileBuilder
	WithType(typ string) FileBuilder
	WithHash(h hash.Hash) FileBuilder
	Now() (File, error)
}
