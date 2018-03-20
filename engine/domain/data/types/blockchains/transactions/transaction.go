package domain

import (
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
)

// Transaction represents a Transaction
type Transaction interface {
	GetMetaData() met.MetaData
	GetJSON() []byte
}
