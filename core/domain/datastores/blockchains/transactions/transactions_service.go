package domain

import (
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions"
)

// TransactionsService represents a Transactions service
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_transactions.Transactions, error)
}
