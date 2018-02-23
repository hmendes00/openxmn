package currencies

import uuid "github.com/satori/go.uuid"

// Save represents a save currency transaction
type Save interface {
	GetID() *uuid.UUID
	GetSymbol() string
	GetName() string
	GetDescription() string
}
