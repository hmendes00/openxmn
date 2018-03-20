package domain

import (
	stored_signed_transaction "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
)

// TransactionsService represents a TRansactionsService instance
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_signed_transaction.Transactions, error)
}
