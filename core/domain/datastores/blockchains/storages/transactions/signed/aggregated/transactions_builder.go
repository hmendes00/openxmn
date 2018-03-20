package aggregated

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
)

// TransactionsBuilder represents aggregated transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithMetaData(met stored_files.File) TransactionsBuilder
	WithTransactions(trs stored_signed_transactions.Transactions) TransactionsBuilder
	WithAtomicTransactions(atomicTrs stored_signed_transactions.AtomicTransactions) TransactionsBuilder
	Now() (Transactions, error)
}
