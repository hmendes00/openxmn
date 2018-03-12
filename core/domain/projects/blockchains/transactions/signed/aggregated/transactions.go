package domain

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed"
)

// Transactions represents aggregated signed transactions
type Transactions interface {
	GetMetaData() metadata.MetaData
	HasTransactions() bool
	GetTransactions() signed_transactions.Transactions
	HasAtomicTransactions() bool
	GetAtomicTransactions() signed_transactions.AtomicTransactions
}
