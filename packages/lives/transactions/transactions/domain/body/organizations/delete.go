package organizations

import uuid "github.com/satori/go.uuid"

// Delete represents a transaction used to delete an organization
type Delete interface {
	GetID() *uuid.UUID
}
