package infrastructure

import (
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

type block struct {
	metadata stored_files.File
	ht       stored_files.File
	trs      []stored_aggregated_transactions.SignedTransactions
}

func createBlock(metadata stored_files.File, ht stored_files.File, trs []stored_aggregated_transactions.SignedTransactions) stored_blocks.Block {
	out := block{
		metadata: metadata,
		ht:       ht,
		trs:      trs,
	}

	return &out
}

// GetMetaData returns the metadata file
func (blk *block) GetMetaData() stored_files.File {
	return blk.metadata
}

// GetHashTree returns the hashtree file
func (blk *block) GetHashTree() stored_files.File {
	return blk.ht
}

// GetTransactions returns the stored transactions
func (blk *block) GetTransactions() []stored_aggregated_transactions.SignedTransactions {
	return blk.trs
}
