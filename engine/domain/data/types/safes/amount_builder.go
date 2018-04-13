package safes

import uuid "github.com/satori/go.uuid"

// AmountBuilder represents an amount builder
type AmountBuilder interface {
	Create() AmountBuilder
	WithTokenID(id *uuid.UUID) AmountBuilder
	WithAmount(amount float64) AmountBuilder
	Now() (Amount, error)
}
