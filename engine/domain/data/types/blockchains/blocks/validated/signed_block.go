package domain

import (
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
)

// SignedBlock represents a signed block of transactions
type SignedBlock interface {
	GetMetaData() metadata.MetaData
	GetBlock() Block
	GetSignature() users.Signature
}
