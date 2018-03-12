package domain

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
)

// TransactionsService represents a transactions service
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_aggregated_transactions.Transactions, error)
}
