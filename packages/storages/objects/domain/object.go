package domain

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Object represents a stored object
type Object interface {
	GetMetaData() stored_files.File
	HasChunks() bool
	GetChunks() stored_chunks.Chunks
}
