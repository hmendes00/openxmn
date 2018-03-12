package domain

import (
	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
)

// User represents a container of coins
type User interface {
	GetMetaData() metadata.MetaData
	GetPublicKey() cryptography.PublicKey
}
