package cities

import uuid "github.com/satori/go.uuid"

// Delete represents a delete city transaction
type Delete interface {
	GetID() *uuid.UUID
}
