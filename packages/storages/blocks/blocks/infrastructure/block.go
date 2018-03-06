package infrastructure

import (
	stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

type block struct {
	met stored_files.File
	trs []stored_aggregated_transactions.SignedTransactions
}

func createBlock(met stored_files.File, trs []stored_aggregated_transactions.SignedTransactions) stored_blocks.Block {
	out := block{
		met: met,
		trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *block) GetMetaData() stored_files.File {
	return trs.met
}

// GetTransactions returns the SignedTransactions
func (trs *block) GetTransactions() []stored_aggregated_transactions.SignedTransactions {
	return trs.trs
}
