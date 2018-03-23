package domain

import (
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// Transaction represents a signed transaction
type Transaction interface {
	GetMetaData() met.MetaData
	GetTransaction() transactions.Transaction
	GetSignature() users.Signature
}
