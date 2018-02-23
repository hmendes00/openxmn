package countries

import uuid "github.com/satori/go.uuid"

// Delete represents a delete country transaction
type Delete interface {
	GetID() *uuid.UUID
}
