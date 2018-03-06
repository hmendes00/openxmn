package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// AtomicTransactionsBuilder represents stored signed AtomicTransactionsBuilder
type AtomicTransactionsBuilder interface {
	Create() AtomicTransactionsBuilder
	WithMetaData(met stored_files.File) AtomicTransactionsBuilder
	WithTransactions(trs []AtomicTransaction) AtomicTransactionsBuilder
	Now() (AtomicTransactions, error)
}
