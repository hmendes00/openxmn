package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignedBlock represents a stored signed block
type SignedBlock interface {
	GetMetaData() stored_files.File
	GetSignature() stored_files.File
	GetBlock() Block
}
