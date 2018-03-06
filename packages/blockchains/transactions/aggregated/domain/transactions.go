package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
)

// Transactions represents aggregated signed transactions
type Transactions interface {
	GetMetaData() metadata.MetaData
	HasTransactions() bool
	GetTransactions() signed_transactions.Transactions
	HasAtomicTransactions() bool
	GetAtomicTransactions() signed_transactions.AtomicTransactions
}
