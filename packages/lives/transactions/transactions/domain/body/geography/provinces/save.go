package provinces

import uuid "github.com/satori/go.uuid"

// Save represents a save province transaction
type Save interface {
	GetID() *uuid.UUID
	GetCode() string
	GetName() string
}
