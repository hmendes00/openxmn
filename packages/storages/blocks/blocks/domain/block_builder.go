package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// BlockBuilder represents a stored block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithHashTree(ht stored_files.File) BlockBuilder
	WithTransactions(trs []stored_aggregated_transactions.SignedTransactions) BlockBuilder
	Now() (Block, error)
}
