package domain

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
)

// Transactions represents signed []Transaction ordered by an HashMap
type Transactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []Transaction
}
