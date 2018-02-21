package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignedTransactions represents aggregated signed transactions
type SignedTransactions interface {
	GetMetaData() stored_files.File
	GetSignature() stored_files.File
	GetTransactions() Transactions
}
