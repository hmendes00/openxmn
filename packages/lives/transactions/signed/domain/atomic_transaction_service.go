package domain

import (
	stored_signed_transaction "github.com/XMNBlockchain/core/packages/storages/transactions/signed/domain"
)

// AtomicTransactionService represents a transaction service
type AtomicTransactionService interface {
	Save(dirPath string, trs AtomicTransaction) (stored_signed_transaction.AtomicTransaction, error)
	SaveAll(dirPath string, trs []AtomicTransaction) ([]stored_signed_transaction.AtomicTransaction, error)
}
