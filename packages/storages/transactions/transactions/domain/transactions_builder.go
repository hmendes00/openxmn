package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// TransactionsBuilder represents a Transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithMetaData(met stored_files.File) TransactionsBuilder
	WithTransactions(trs []Transaction) TransactionsBuilder
	Now() (Transactions, error)
}
