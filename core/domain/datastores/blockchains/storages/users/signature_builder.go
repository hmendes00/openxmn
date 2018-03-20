package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
)

// SignatureBuilder represents a stored signature builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithMetaData(met stored_files.File) SignatureBuilder
	WithSignature(sig stored_files.File) SignatureBuilder
	WithUser(usr User) SignatureBuilder
	Now() (Signature, error)
}
