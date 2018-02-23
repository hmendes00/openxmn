package branches

import uuid "github.com/satori/go.uuid"

// Delete represents a delete branch transaction
type Delete interface {
	GetID() *uuid.UUID
}
