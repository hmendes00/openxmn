package domain

import (
	stored_signed_transaction "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// TransactionsService represents a TRansactionsService instance
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_signed_transaction.Transactions, error)
}
