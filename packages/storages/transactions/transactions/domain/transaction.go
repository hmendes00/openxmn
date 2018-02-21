package domain

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Transaction represents a stored transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetChunks() stored_chunks.Chunks
}
