package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
)

// AtomicTransactions represents signed []AtomicTransaction ordered by an HashMap
type AtomicTransactions interface {
	GetMetaData() metadata.MetaData
	GetTransactions() []AtomicTransaction
}
