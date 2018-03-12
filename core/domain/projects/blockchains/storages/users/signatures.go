package domain

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// Signatures represents stored signatures
type Signatures interface {
	GetMetaData() stored_files.File
	GetSignatures() []Signature
}
