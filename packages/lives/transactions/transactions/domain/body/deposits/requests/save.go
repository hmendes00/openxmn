package requests

import uuid "github.com/satori/go.uuid"

// Save represents a save transaction request
type Save interface {
	GetID() *uuid.UUID
	GetAccountID() *uuid.UUID
	GetUserID() *uuid.UUID
	GetAmount() float64
}
