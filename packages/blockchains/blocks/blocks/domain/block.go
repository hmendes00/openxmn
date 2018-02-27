package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	uuid "github.com/satori/go.uuid"
)

// Block represents a Block instance
type Block interface {
	GetID() *uuid.UUID
	GetHashTree() hashtrees.HashTree
	GetTransactions() []aggregated.SignedTransactions
	CreatedOn() time.Time
}
