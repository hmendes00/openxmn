package infrastructure

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
)

// AggregatedSignedTransactions represents a concrete AggregatedSignedTransactions implementation
type AggregatedSignedTransactions struct {
	Met *concrete_metadata.MetaData `json:"metadata"`
	Trs []*SignedTransactions       `json:"signed_transactions"`
}

func createAggregatedSignedTransactions(met *concrete_metadata.MetaData, trs []*SignedTransactions) aggregated.AggregatedSignedTransactions {
	out := AggregatedSignedTransactions{
		Met: met,
		Trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *AggregatedSignedTransactions) GetMetaData() metadata.MetaData {
	return trs.Met
}

// GetTransactions returns the []SignedTransactions
func (trs *AggregatedSignedTransactions) GetTransactions() []aggregated.SignedTransactions {
	out := []aggregated.SignedTransactions{}
	for _, oneTrs := range trs.Trs {
		out = append(out, oneTrs)
	}

	return out
}
