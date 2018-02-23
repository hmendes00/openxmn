package streets

import uuid "github.com/satori/go.uuid"

// Delete represents a delete street transaction
type Delete interface {
	GetID() *uuid.UUID
}
