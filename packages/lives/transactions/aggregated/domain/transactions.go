package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	uuid "github.com/satori/go.uuid"
)

// Transactions represents aggregated signed transactions
type Transactions interface {
	GetID() *uuid.UUID
	GetHashTree() hashtrees.HashTree
	HasTrs() bool
	GetTrs() []signed_transactions.Transaction
	HasAtomicTrs() bool
	GetAtomicTrs() []signed_transactions.AtomicTransaction
	CreatedOn() time.Time
}
