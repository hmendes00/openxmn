package domain

import (
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// TransactionsService represents a Transactions service
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_transactions.Transactions, error)
}
