package domain

import (
	stored_signed_transaction "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed"
)

// TransactionsService represents a TRansactionsService instance
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_signed_transaction.Transactions, error)
}
