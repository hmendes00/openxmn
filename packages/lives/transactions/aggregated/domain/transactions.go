package domain

import (
	"time"

	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	uuid "github.com/satori/go.uuid"
)

// Transactions represents aggregated signed transactions
type Transactions interface {
	GetID() *uuid.UUID
	GetTrs() []signed_transactions.Transaction
	GetAtomicTrs() []signed_transactions.AtomicTransaction
	CreatedOn() time.Time
}
