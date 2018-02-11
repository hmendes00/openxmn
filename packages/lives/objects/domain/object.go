package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	chunks "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// Object represents an object
type Object interface {
	GetHashTree() hashtrees.HashTree
	GetID() *uuid.UUID
	GetPath() string
	CreatedOn() time.Time
	HasSignature() bool
	GetSignature() users.Signature
	HasChunks() bool
	GetChunks() chunks.Chunks
}
