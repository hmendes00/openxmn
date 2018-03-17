package users

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// User represents a stored user
type User interface {
	GetMetaData() stored_files.File
	GetPublicKey() stored_files.File
}
