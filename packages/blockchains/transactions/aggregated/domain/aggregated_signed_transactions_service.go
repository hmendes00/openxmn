package domain

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// AggregatedSignedTransactionsService represents an aggregated signed transactions service
type AggregatedSignedTransactionsService interface {
	Save(dirPath string, trs AggregatedSignedTransactions) (stored_aggregated_transactions.AggregatedSignedTransactions, error)
}
