package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed"
)

// Transactions represents aggregated transactions
type Transactions interface {
	GetMetaData() stored_files.File
	HasTransactions() bool
	GetTransactions() stored_signed_transactions.Transactions
	HasAtomicTransactions() bool
	GetAtomicTransactions() stored_signed_transactions.AtomicTransactions
}
