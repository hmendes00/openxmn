package domain

import (
	stored_signed_transaction "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
)

// AtomicTransactionsService represents an AtomicTransactionsService instance
type AtomicTransactionsService interface {
	Save(dirPath string, trs AtomicTransactions) (stored_signed_transaction.AtomicTransactions, error)
}
