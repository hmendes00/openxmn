package blockchains

import uuid "github.com/satori/go.uuid"

// Delete represents a delete blockchain transaction
type Delete interface {
	GetID() *uuid.UUID
}
