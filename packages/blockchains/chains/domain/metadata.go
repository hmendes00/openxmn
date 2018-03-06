package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents the chain metadata
type MetaData interface {
	GetID() *uuid.UUID
	GetHashTree() hashtrees.HashTree
	GetHeight() int
	CreatedOn() time.Time
	LastUpdatedOn() time.Time
}
