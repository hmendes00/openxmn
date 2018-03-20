package blocks

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
)

// BlockBuilder represents a stored block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithTransactions(trs []stored_aggregated_transactions.SignedTransactions) BlockBuilder
	Now() (Block, error)
}
