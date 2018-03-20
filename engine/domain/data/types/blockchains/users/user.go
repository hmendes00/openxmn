package domain

import (
	cryptography "github.com/XMNBlockchain/exmachina-network/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/metadata"
)

// User represents a container of coins
type User interface {
	GetMetaData() metadata.MetaData
	GetPublicKey() cryptography.PublicKey
}
