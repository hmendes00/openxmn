package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Signature represents a stored Signature
type Signature interface {
	GetMetaData() stored_files.File
	GetSignature() stored_files.File
	GetUser() User
}
