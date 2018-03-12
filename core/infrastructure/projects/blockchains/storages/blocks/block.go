package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
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
