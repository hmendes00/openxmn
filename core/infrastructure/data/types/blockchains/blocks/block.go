package blocks

import (
	blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/blocks"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	concrete_aggregated "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions/signed/aggregated"
)

// Block represents a concrete Block implementation
type Block struct {
	Met *concrete_metadata.MetaData               `json:"metadata"`
	Trs []*concrete_aggregated.SignedTransactions `json:"signed_transactions"`
}

func createBlock(met *concrete_metadata.MetaData, trs []*concrete_aggregated.SignedTransactions) blocks.Block {
	out := Block{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *Block) GetMetaData() metadata.MetaData {
	return trs.Met
}

// GetTransactions returns the []SignedTransactions
func (trs *Block) GetTransactions() []aggregated.SignedTransactions {
	out := []aggregated.SignedTransactions{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}

	return out
}
