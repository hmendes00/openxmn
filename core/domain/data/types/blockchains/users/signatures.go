package domain

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
)

// Signatures represents []Signature ordered by an HashMap
type Signatures interface {
	GetMetaData() metadata.MetaData
	GetSignatures() []Signature
}