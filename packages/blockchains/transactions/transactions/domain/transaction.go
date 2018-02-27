package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// Transaction represents a Transaction
type Transaction interface {
	GetID() *uuid.UUID
	GetJSON() []byte
	CreatedOn() time.Time
}
