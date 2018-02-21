package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// TransactionBuilder represents a stored signed transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithMetaData(met stored_files.File) TransactionBuilder
	WithSignature(sig stored_files.File) TransactionBuilder
	WithTransaction(trs stored_transactions.Transaction) TransactionBuilder
	Now() (Transaction, error)
}
