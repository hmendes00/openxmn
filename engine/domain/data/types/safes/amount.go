package safes

import uuid "github.com/satori/go.uuid"

// Amount represents an amount in a safe
type Amount interface {
	GetTokenID() *uuid.UUID
	GetAmount() float64
}
