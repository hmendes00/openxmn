package blocks

import (
	stored_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed/aggregated"
	conrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
	concrete_stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/transactions/signed/aggregated"
)

// Block represents a concrete stored block implementation
type Block struct {
	Met *conrete_stored_files.File                                    `json:"metadata"`
	Trs []*concrete_stored_aggregated_transactions.SignedTransactions `json:"aggregated_signed_transactions"`
}

func createBlock(met *conrete_stored_files.File, trs []*concrete_stored_aggregated_transactions.SignedTransactions) stored_blocks.Block {
	out := Block{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *Block) GetMetaData() stored_files.File {
	return trs.Met
}

// GetTransactions returns the SignedTransactions
func (trs *Block) GetTransactions() []stored_aggregated_transactions.SignedTransactions {
	out := []stored_aggregated_transactions.SignedTransactions{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}
	return out
}
