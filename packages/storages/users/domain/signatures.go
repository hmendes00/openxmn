package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Signatures represents stored signatures
type Signatures interface {
	GetMetaData() stored_files.File
	GetSignatures() []Signature
}
