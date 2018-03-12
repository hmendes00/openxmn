package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/hashtrees"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents a chained block metadata
type MetaData interface {
	GetID() *uuid.UUID
	GetPreviousID() *uuid.UUID
	GetHashTree() hashtrees.HashTree
	CreatedOn() time.Time
}
