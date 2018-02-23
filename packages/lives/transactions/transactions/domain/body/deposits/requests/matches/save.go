package matches

import uuid "github.com/satori/go.uuid"

// Save represents a save matched transaction request
type Save interface {
	GetID() *uuid.UUID
	GetRequestID() *uuid.UUID
	GetAmount() int
}
