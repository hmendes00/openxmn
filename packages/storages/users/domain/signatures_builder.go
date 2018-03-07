package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignaturesBuilder represents a Signatures builder
type SignaturesBuilder interface {
	Create() SignaturesBuilder
	WithMetaData(met stored_files.File) SignaturesBuilder
	WithSignatures(sigs []Signature) SignaturesBuilder
	Now() (Signatures, error)
}
