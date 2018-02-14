package domain

import (
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
)

// ObjectBuilder represents an ObjectBuilder
type ObjectBuilder interface {
	Create() ObjectBuilder
	WithMetaData(met MetaData) ObjectBuilder
	WithChunks(chks chunks.Chunks) ObjectBuilder
	Now() (Object, error)
}
