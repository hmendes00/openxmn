package domain

import (
	"time"

	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// TransactionBuilder represents a signed transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithID(id *uuid.UUID) TransactionBuilder
	WithTransaction(trs transactions.Transaction) TransactionBuilder
	WithSignature(sig users.Signature) TransactionBuilder
	CreatedOn(ts time.Time) TransactionBuilder
	Now() (Transaction, error)
}
