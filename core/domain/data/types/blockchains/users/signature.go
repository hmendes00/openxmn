package domain

import (
	cryptography "github.com/XMNBlockchain/exmachina-network/core/domain/cryptography"
	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
)

// Signature represents the Signature of a User
type Signature interface {
	GetMetaData() metadata.MetaData
	GetSignature() cryptography.Signature
	GetUser() User
	String() string
}
