package domain

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// Transactions represents signed []Transaction ordered by an HashMap
type Transactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []Transaction
}
