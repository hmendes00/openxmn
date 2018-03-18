package domain

import (
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/users"
)

// SignedBlock represents a SignedBlock instance
type SignedBlock interface {
	GetMetaData() metadata.MetaData
	GetBlock() Block
	GetSignature() users.Signature
}