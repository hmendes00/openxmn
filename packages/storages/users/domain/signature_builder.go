package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignatureBuilder represents a stored signature builder
type SignatureBuilder interface {
	Create() SignatureBuilder
	WithMetaData(met stored_files.File) SignatureBuilder
	WithSignature(sig stored_files.File) SignatureBuilder
	WithUser(usr User) SignatureBuilder
	Now() (Signature, error)
}
