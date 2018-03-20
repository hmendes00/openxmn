package blocks

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
)

// Block represents a stored block
type Block interface {
	GetMetaData() stored_files.File
	GetTransactions() []stored_aggregated_transactions.SignedTransactions
}
