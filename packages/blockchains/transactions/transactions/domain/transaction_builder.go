package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// TransactionBuilder represents the Transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithID(id *uuid.UUID) TransactionBuilder
	WithJSON(data []byte) TransactionBuilder
	CreatedOn(time time.Time) TransactionBuilder
	Now() (Transaction, error)
}
