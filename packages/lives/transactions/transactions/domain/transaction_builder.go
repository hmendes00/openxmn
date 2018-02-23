package domain

import (
	"time"

	body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain/body"
	uuid "github.com/satori/go.uuid"
)

// TransactionBuilder represents the Transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithID(id *uuid.UUID) TransactionBuilder
	WithBody(bod body.Body) TransactionBuilder
	CreatedOn(time time.Time) TransactionBuilder
	Now() (Transaction, error)
}
