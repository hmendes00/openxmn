package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// AtomicTransactionBuilder represents a signed atomic transaction builder
type AtomicTransactionBuilder interface {
	Create() AtomicTransactionBuilder
	WithMetaData(met stored_files.File) AtomicTransactionBuilder
	WithSignature(sig stored_users.Signature) AtomicTransactionBuilder
	WithTransactions(trs stored_transactions.Transactions) AtomicTransactionBuilder
	Now() (AtomicTransaction, error)
}
