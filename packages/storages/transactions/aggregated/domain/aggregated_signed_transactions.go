package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// AggregatedSignedTransactions represents aggregated aggregated signed transactions
type AggregatedSignedTransactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []SignedTransactions
}
