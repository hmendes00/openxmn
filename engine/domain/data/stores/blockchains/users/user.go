package users

import (
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// User represents a stored user
type User interface {
	GetMetaData() stored_files.File
	GetPublicKey() stored_files.File
}
