package tokens

import uuid "github.com/satori/go.uuid"

// Delete represents a transaction used to delete a token
type Delete interface {
	GetID() *uuid.UUID
}
