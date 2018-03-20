package domain

import (
	stored_signed_transaction "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
)

// TransactionService represents a transaction service
type TransactionService interface {
	Save(dirPath string, trs Transaction) (stored_signed_transaction.Transaction, error)
	SaveAll(dirPath string, trs []Transaction) ([]stored_signed_transaction.Transaction, error)
}
