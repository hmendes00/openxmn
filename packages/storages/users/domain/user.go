package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// User represents a stored user
type User interface {
	GetMetaData() stored_files.File
	GetPublicKey() stored_files.File
}
