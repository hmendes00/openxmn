package blocks

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed/aggregated"
)

// BlockBuilder represents a stored block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithTransactions(trs []stored_aggregated_transactions.SignedTransactions) BlockBuilder
	Now() (Block, error)
}
