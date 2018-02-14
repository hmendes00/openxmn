package domain

import (
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
)

// Object represents an object
type Object interface {
	GetMetaData() MetaData
	HasChunks() bool
	GetChunks() chunks.Chunks
}
