package transactions

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Transactions represents a stored Transactions
type Transactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []Transaction
}
