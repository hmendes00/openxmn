package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// TransactionsBuilder represents a Transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithMetaData(met stored_files.File) TransactionsBuilder
	WithTransactions(trs []Transaction) TransactionsBuilder
	Now() (Transactions, error)
}
