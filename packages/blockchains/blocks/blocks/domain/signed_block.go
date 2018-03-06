package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
)

// SignedBlock represents a SignedBlock instance
type SignedBlock interface {
	GetMetaData() metadata.MetaData
	GetBlock() Block
	GetSignature() users.Signature
}
