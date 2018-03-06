package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// AggregatedSignedTransactionsBuilder represents aggregated aggregated signed transactions builder
type AggregatedSignedTransactionsBuilder interface {
	Create() AggregatedSignedTransactionsBuilder
	WithMetaData(met stored_files.File) AggregatedSignedTransactionsBuilder
	WithTransactions(trs []SignedTransactions) AggregatedSignedTransactionsBuilder
	Now() (AggregatedSignedTransactions, error)
}
