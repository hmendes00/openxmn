package domain

import (
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	cryptography "github.com/XMNBlockchain/core/packages/cryptography/domain"
)

// Signature represents the Signature of a User
type Signature interface {
	GetMetaData() metadata.MetaData
	GetSignature() cryptography.Signature
	GetUser() User
}
