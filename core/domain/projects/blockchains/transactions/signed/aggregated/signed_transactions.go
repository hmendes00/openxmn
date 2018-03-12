package domain

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
)

// SignedTransactions represents the aggregated signed transactions, then signed by a pointer
type SignedTransactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() Transactions
	GetSignature() users.Signature
}
