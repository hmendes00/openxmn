package domain

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	hashtree "github.com/XMNBlockchain/core/packages/lives/hashtrees/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	uuid "github.com/satori/go.uuid"
)

// Block represents a block of transactions
type Block interface {
	GetID() *uuid.UUID
	GetHashTree() hashtree.HashTree
	GetBlock() blocks.SignedBlock
	GetSignatures() []users.Signature
	CreatedOn() time.Time
}
