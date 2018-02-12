package domain

import (
	"time"

	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	uuid "github.com/satori/go.uuid"
)

// ObjectBuilder represents a stored ObjectBuilder
type ObjectBuilder interface {
	Create() ObjectBuilder
	WithID(id *uuid.UUID) ObjectBuilder
	WithSignature(sig stored_files.File) ObjectBuilder
	WithChunks(chks stored_chunks.Chunks) ObjectBuilder
	CreatedOn(ts time.Time) ObjectBuilder
	Now() (Object, error)
}
