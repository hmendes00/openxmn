package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
)

// Transactions represents []Transaction ordered by an HashMap
type Transactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []Transaction
}
