package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/files"
)

// SignaturesBuilder represents a Signatures builder
type SignaturesBuilder interface {
	Create() SignaturesBuilder
	WithMetaData(met stored_files.File) SignaturesBuilder
	WithSignatures(sigs []Signature) SignaturesBuilder
	Now() (Signatures, error)
}
