package domain

import (
	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// Transaction represents a signed transaction
type Transaction interface {
	GetMetaData() met.MetaData
	GetTransaction() transactions.Transaction
	GetSignature() users.Signature
}
