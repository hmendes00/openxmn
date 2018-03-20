package signed

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Transactions represents stored signed Transactions
type Transactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []Transaction
}
