package domain

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// Block represents a block of transactions
type Block interface {
	GetID() *uuid.UUID
	GetBlock() blocks.SignedBlock
	GetSignatures() []users.Signature
	CreatedOn() time.Time
}
