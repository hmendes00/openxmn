package currencies

import uuid "github.com/satori/go.uuid"

// Delete represents a delete currency transaction
type Delete interface {
	GetID() *uuid.UUID
}
