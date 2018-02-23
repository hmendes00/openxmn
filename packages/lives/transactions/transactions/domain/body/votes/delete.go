package votes

import uuid "github.com/satori/go.uuid"

// Delete represents a delete vote transaction
type Delete interface {
	GetID() *uuid.UUID
}
