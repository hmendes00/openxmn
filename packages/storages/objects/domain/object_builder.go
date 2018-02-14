package domain

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// ObjectBuilder represents a stored ObjectBuilder
type ObjectBuilder interface {
	Create() ObjectBuilder
	WithMetaData(metaData stored_files.File) ObjectBuilder
	WithChunks(chks stored_chunks.Chunks) ObjectBuilder
	Now() (Object, error)
}
