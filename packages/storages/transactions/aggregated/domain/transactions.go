package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// Transactions represents aggregated transactions
type Transactions interface {
	GetMetaData() stored_files.File
	GetHashTree() stored_files.File
	HasTrs() bool
	GetTrs() []stored_signed_transactions.Transaction
	HasAtomicTrs() bool
	GetAtomicTrs() []stored_signed_transactions.AtomicTransaction
}
