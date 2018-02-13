package domain

import (
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	uuid "github.com/satori/go.uuid"
)

// Object represents a stored object
type Object interface {
	GetID() *uuid.UUID
	CreatedOn() time.Time
	HasSignature() bool
	GetSignature() stored_files.File
	HasChunks() bool
	GetChunks() stored_chunks.Chunks
}
