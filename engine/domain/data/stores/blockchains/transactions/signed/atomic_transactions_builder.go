package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
)

// AtomicTransactionsBuilder represents stored signed AtomicTransactionsBuilder
type AtomicTransactionsBuilder interface {
	Create() AtomicTransactionsBuilder
	WithMetaData(met stored_files.File) AtomicTransactionsBuilder
	WithTransactions(trs []AtomicTransaction) AtomicTransactionsBuilder
	Now() (AtomicTransactions, error)
}