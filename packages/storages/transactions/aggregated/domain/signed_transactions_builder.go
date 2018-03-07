package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	stored_users "github.com/XMNBlockchain/core/packages/storages/users/domain"
)

// SignedTransactionsBuilder represents aggregated signed transactions builder
type SignedTransactionsBuilder interface {
	Create() SignedTransactionsBuilder
	WithMetaData(met stored_files.File) SignedTransactionsBuilder
	WithSignature(sig stored_users.Signature) SignedTransactionsBuilder
	WithTransactions(trs Transactions) SignedTransactionsBuilder
	Now() (SignedTransactions, error)
}
