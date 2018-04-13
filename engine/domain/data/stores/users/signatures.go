package users

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Signatures represents stored signatures
type Signatures interface {
	GetMetaData() stored_files.File
	GetSignatures() []Signature
}
