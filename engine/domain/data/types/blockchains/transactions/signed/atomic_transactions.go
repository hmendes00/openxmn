package domain

import (
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
)

// AtomicTransactions represents signed []AtomicTransaction ordered by an HashMap
type AtomicTransactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []AtomicTransaction
}
