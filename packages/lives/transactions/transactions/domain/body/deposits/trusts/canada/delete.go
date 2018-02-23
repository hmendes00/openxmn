package canada

import uuid "github.com/satori/go.uuid"

// Delete represents a delete canadian trust transaction
type Delete interface {
	GetID() *uuid.UUID
}
