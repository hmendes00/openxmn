package servers

import uuid "github.com/satori/go.uuid"

// Delete represents a transaction used to delete a server
type Delete interface {
	GetID() *uuid.UUID
}
