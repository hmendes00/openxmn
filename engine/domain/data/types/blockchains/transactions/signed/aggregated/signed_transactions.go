package domain

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/users"
)

// SignedTransactions represents the aggregated signed transactions, then signed by a pointer
type SignedTransactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() Transactions
	GetSignature() users.Signature
}
