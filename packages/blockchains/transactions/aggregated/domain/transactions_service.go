package domain

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
)

// TransactionsService represents a transactions service
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_aggregated_transactions.Transactions, error)
}
