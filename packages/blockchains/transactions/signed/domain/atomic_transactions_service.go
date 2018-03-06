package domain

import (
	stored_signed_transaction "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// AtomicTransactionsService represents an AtomicTransactionsService instance
type AtomicTransactionsService interface {
	Save(dirPath string, trs AtomicTransactions) (stored_signed_transaction.AtomicTransactions, error)
}
