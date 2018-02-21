package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
)

// AtomicTransactionBuilder represents a signed atomic transaction builder
type AtomicTransactionBuilder interface {
	Create() AtomicTransactionBuilder
	WithMetaData(met stored_files.File) AtomicTransactionBuilder
	WithSignature(sig stored_files.File) AtomicTransactionBuilder
	WithHashTree(ht stored_files.File) AtomicTransactionBuilder
	WithTransactions(trs []stored_transactions.Transaction) AtomicTransactionBuilder
	Now() (AtomicTransaction, error)
}
