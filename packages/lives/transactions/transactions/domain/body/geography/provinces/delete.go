package provinces

import uuid "github.com/satori/go.uuid"

// Delete represents a delete province transaction
type Delete interface {
	GetID() *uuid.UUID
}
