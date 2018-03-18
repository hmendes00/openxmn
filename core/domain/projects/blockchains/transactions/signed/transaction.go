package domain

import (
	met "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
)

// Transaction represents a signed transaction
type Transaction interface {
	GetMetaData() met.MetaData
	GetTransaction() transactions.Transaction
	GetSignature() users.Signature
}