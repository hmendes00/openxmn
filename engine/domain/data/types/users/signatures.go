package domain

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
)

// Signatures represents []Signature ordered by an HashMap
type Signatures interface {
	GetMetaData() metadata.MetaData
	GetSignatures() []Signature
}
