package applications

import uuid "github.com/satori/go.uuid"

// Delete represents a delete application transaction
type Delete interface {
	GetID() *uuid.UUID
}
