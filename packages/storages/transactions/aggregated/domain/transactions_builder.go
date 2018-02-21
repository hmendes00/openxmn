package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// TransactionsBuilder represents aggregated transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithMetaData(met stored_files.File) TransactionsBuilder
	WithHashTree(ht stored_files.File) TransactionsBuilder
	WithTrs(trs []stored_signed_transactions.Transaction) TransactionsBuilder
	WithAtomicTrs(atomicTrs []stored_signed_transactions.AtomicTransaction) TransactionsBuilder
	Now() (Transactions, error)
}
