package accounts

import uuid "github.com/satori/go.uuid"

// Delete represents a delete account transaction
type Delete interface {
	GetID() *uuid.UUID
}
