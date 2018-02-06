package users

import uuid "github.com/satori/go.uuid"

// Delete represents a transaction used to delete a user
type Delete interface {
	GetID() *uuid.UUID
}
