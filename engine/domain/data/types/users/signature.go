package domain

import (
	cryptography "github.com/XMNBlockchain/openxmn/engine/domain/cryptography"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
)

// Signature represents the Signature of a User
type Signature interface {
	GetMetaData() metadata.MetaData
	GetSignature() cryptography.Signature
	GetUser() User
	String() string
}
