package domain

import (
	stored_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions"
)

// TransactionService represents a transaction service
type TransactionService interface {
	Save(dirPath string, trs Transaction) (stored_transactions.Transaction, error)
	SaveAll(dirPath string, trs []Transaction) ([]stored_transactions.Transaction, error)
}