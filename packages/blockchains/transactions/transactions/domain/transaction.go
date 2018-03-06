package domain

import (
	met "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
)

// Transaction represents a Transaction
type Transaction interface {
	GetMetaData() met.MetaData
	GetJSON() []byte
}
