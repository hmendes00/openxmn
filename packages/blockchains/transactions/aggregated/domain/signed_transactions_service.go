package domain

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// SignedTransactionsService represents a signed transactions service
type SignedTransactionsService interface {
	Save(dirPath string, trs SignedTransactions) (stored_aggregated_transactions.SignedTransactions, error)
	SaveAll(dirPath string, trs []SignedTransactions) ([]stored_aggregated_transactions.SignedTransactions, error)
}
