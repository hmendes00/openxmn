package domain

import (
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
)

// TransactionsService represents a Transactions service
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_transactions.Transactions, error)
}
