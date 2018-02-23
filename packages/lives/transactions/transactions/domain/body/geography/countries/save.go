package countries

import uuid "github.com/satori/go.uuid"

// Save represents a save country transaction
type Save interface {
	GetID() *uuid.UUID
	GetCode() string
	GetName() string
}
