package requests

import uuid "github.com/satori/go.uuid"

// Delete represents a delete request transaction
type Delete interface {
	GetID() *uuid.UUID
}
