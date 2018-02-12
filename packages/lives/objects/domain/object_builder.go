package domain

import (
	"time"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// ObjectBuilder represents an ObjectBuilder
type ObjectBuilder interface {
	Create() ObjectBuilder
	WithID(id *uuid.UUID) ObjectBuilder
	WithSignature(sig users.Signature) ObjectBuilder
	WithChunks(chks chunks.Chunks) ObjectBuilder
	CreatedOn(ts time.Time) ObjectBuilder
	Now() (Object, error)
}
