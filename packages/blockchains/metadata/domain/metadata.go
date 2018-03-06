package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents an object metadata
type MetaData interface {
	GetID() *uuid.UUID
	GetHashTree() hashtrees.HashTree
	CreatedOn() time.Time
}
