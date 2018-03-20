package domain

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/transactions/signed/aggregated"
)

// SignedTransactionsService represents a signed transactions service
type SignedTransactionsService interface {
	Save(dirPath string, trs SignedTransactions) (stored_aggregated_transactions.SignedTransactions, error)
	SaveAll(dirPath string, trs []SignedTransactions) ([]stored_aggregated_transactions.SignedTransactions, error)
}
