package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Transactions represents stored signed Transactions
type Transactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []Transaction
}
