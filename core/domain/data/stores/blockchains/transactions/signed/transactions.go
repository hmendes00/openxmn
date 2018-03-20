package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
)

// Transactions represents stored signed Transactions
type Transactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []Transaction
}
