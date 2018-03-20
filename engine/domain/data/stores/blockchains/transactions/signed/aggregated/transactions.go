package aggregated

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
)

// Transactions represents aggregated transactions
type Transactions interface {
	GetMetaData() stored_files.File
	HasTransactions() bool
	GetTransactions() stored_signed_transactions.Transactions
	HasAtomicTransactions() bool
	GetAtomicTransactions() stored_signed_transactions.AtomicTransactions
}
