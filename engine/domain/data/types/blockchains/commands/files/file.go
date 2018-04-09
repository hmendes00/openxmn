package files

import (
	uuid "github.com/satori/go.uuid"
)

// File represents a file in a command
type File interface {
	GetID() *uuid.UUID
	GetType() string
	GetHash() string
}
