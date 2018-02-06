package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	storages_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// Transactions represents stored aggregated transactions
type Transactions interface {
	GetHashTree() hashtrees.HashTree
	GetID() *uuid.UUID
	GetSignature() users.Signature
	GetTrs() []storages_transactions.Transaction
	GetAtomicTrs() []storages_transactions.AtomicTransaction
	CreatedOn() time.Time
}
