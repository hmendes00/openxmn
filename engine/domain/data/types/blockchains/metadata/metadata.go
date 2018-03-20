package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/hashtrees"
	uuid "github.com/satori/go.uuid"
)

// MetaData represents an object metadata
type MetaData interface {
	GetID() *uuid.UUID
	GetHashTree() hashtrees.HashTree
	CreatedOn() time.Time
}
