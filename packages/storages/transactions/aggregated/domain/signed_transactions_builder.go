package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignedTransactionsBuilder represents aggregated signed transactions builder
type SignedTransactionsBuilder interface {
	Create() SignedTransactionsBuilder
	WithMetaData(met stored_files.File) SignedTransactionsBuilder
	WithSignature(sig stored_files.File) SignedTransactionsBuilder
	WithTransactions(trs Transactions) SignedTransactionsBuilder
	Now() (SignedTransactions, error)
}
