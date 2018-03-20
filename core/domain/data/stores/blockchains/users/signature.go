package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/files"
)

// Signature represents a stored Signature
type Signature interface {
	GetMetaData() stored_files.File
	GetSignature() stored_files.File
	GetUser() User
}