package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// TransactionsBuilder represents aggregated transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithMetaData(met stored_files.File) TransactionsBuilder
	WithTransactions(trs stored_signed_transactions.Transactions) TransactionsBuilder
	WithAtomicTransactions(atomicTrs stored_signed_transactions.AtomicTransactions) TransactionsBuilder
	Now() (Transactions, error)
}
