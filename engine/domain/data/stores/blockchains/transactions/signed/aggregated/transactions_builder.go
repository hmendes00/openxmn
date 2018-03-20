package aggregated

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
)

// TransactionsBuilder represents aggregated transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithMetaData(met stored_files.File) TransactionsBuilder
	WithTransactions(trs stored_signed_transactions.Transactions) TransactionsBuilder
	WithAtomicTransactions(atomicTrs stored_signed_transactions.AtomicTransactions) TransactionsBuilder
	Now() (Transactions, error)
}
