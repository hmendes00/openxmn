package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// SignedTransactions represents the aggregated signed transactions, then signed by a pointer
type SignedTransactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() Transactions
	GetSignature() users.Signature
}
