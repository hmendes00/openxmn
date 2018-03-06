package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// Transactions represents aggregated transactions
type Transactions interface {
	GetMetaData() stored_files.File
	HasTransactions() bool
	GetTransactions() stored_signed_transactions.Transactions
	HasAtomicTransactions() bool
	GetAtomicTransactions() stored_signed_transactions.AtomicTransactions
}
