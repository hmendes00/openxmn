package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
)

// Signatures represents []Signature ordered by an HashMap
type Signatures interface {
	GetMetaData() metadata.MetaData
	GetSignatures() []Signature
}
