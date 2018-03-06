package infrastructure

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

type aggregatedSignedTransactions struct {
	met stored_files.File
	trs []stored_aggregated_transactions.SignedTransactions
}

func createAggregatedSignedTransactions(met stored_files.File, trs []stored_aggregated_transactions.SignedTransactions) stored_aggregated_transactions.AggregatedSignedTransactions {
	out := aggregatedSignedTransactions{
		met: met,
		trs: trs,
	}

	return &out
}

// GetMetaData returns the MetaData
func (trs *aggregatedSignedTransactions) GetMetaData() stored_files.File {
	return trs.met
}

// GetTransactions returns the SignedTransactions
func (trs *aggregatedSignedTransactions) GetTransactions() []stored_aggregated_transactions.SignedTransactions {
	return trs.trs
}
