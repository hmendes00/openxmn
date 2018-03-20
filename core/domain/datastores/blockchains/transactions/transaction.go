package domain

import (
	met "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
)

// Transaction represents a Transaction
type Transaction interface {
	GetMetaData() met.MetaData
	GetJSON() []byte
}
