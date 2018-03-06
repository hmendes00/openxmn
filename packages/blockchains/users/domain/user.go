package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// User represents a container of coins
type User interface {
	GetMetaData() metadata.MetaData
	GetPublicKey() cryptography.PublicKey
}
