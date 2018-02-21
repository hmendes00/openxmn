package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// Transaction represents a stored signed transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_files.File
	GetTransaction() stored_transactions.Transaction
}
