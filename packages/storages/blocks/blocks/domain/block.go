package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// Block represents a stored block
type Block interface {
	GetMetaData() stored_files.File
	GetHashTree() stored_files.File
	GetTransactions() []stored_aggregated_transactions.SignedTransactions
}
