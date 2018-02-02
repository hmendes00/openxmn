package custom

import uuid "github.com/satori/go.uuid"

// Create represents a create custom transaction
type Create interface {
	GetID() *uuid.UUID
	GetJSON() []byte
}
