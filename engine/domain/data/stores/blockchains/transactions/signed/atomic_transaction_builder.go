package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
	stored_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/transactions"
	stored_users "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/blockchains/users"
)

// AtomicTransactionBuilder represents a signed atomic transaction builder
type AtomicTransactionBuilder interface {
	Create() AtomicTransactionBuilder
	WithMetaData(met stored_files.File) AtomicTransactionBuilder
	WithSignature(sig stored_users.Signature) AtomicTransactionBuilder
	WithTransactions(trs stored_transactions.Transactions) AtomicTransactionBuilder
	Now() (AtomicTransaction, error)
}
