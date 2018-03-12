package domain

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
)

// Transactions represents []Transaction ordered by an HashMap
type Transactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []Transaction
}
