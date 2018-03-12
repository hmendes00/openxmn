package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// Transactions represents a stored Transactions
type Transactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []Transaction
}
