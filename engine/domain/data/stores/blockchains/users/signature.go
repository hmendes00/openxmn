package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
)

// Signature represents a stored Signature
type Signature interface {
	GetMetaData() stored_files.File
	GetSignature() stored_files.File
	GetUser() User
}
