package domain

import (
	"time"

	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// Object represents an object
type Object interface {
	GetID() *uuid.UUID
	CreatedOn() time.Time
	HasSignature() bool
	GetSignature() users.Signature
	HasChunks() bool
	GetChunks() chunks.Chunks
}
