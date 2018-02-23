package votes

import uuid "github.com/satori/go.uuid"

// Save represents a save vote on a transaction, transaction
type Save interface {
	GetID() *uuid.UUID
	GetTransactionID() *uuid.UUID
	IsApproved() bool
}
