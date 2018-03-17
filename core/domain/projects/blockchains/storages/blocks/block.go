package blocks

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
)

// Block represents a stored block
type Block interface {
	GetMetaData() stored_files.File
	GetTransactions() []stored_aggregated_transactions.SignedTransactions
}
