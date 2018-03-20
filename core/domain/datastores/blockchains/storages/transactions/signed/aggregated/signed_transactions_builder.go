package aggregated

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
)

// SignedTransactionsBuilder represents aggregated signed transactions builder
type SignedTransactionsBuilder interface {
	Create() SignedTransactionsBuilder
	WithMetaData(met stored_files.File) SignedTransactionsBuilder
	WithSignature(sig stored_users.Signature) SignedTransactionsBuilder
	WithTransactions(trs Transactions) SignedTransactionsBuilder
	Now() (SignedTransactions, error)
}
