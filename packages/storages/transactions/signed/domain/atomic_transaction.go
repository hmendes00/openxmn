package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// AtomicTransaction represents a signed atomic transaction
type AtomicTransaction interface {
	GetMetaData() stored_files.File
	GetSignature() stored_files.File
	GetHashTree() stored_files.File
	GetTransactions() []stored_transactions.Transaction
}
