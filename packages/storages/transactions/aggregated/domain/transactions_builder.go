package domain

import (
	"time"

	storages_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// TransactionsBuilder represents stored aggregated transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithID(id *uuid.UUID) TransactionsBuilder
	WithSignature(sig users.Signature) TransactionsBuilder
	WithTrs(trs []storages_transactions.Transaction) TransactionsBuilder
	WithAtomicTrs(atomicTrs []storages_transactions.AtomicTransaction) TransactionsBuilder
	CreatedOn(ts time.Time) TransactionsBuilder
	Now() (Transactions, error)
}
