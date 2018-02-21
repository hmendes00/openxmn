package domain

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/lives/blocks/blocks/domain"
	users "github.com/XMNBlockchain/core/packages/lives/users/domain"
	uuid "github.com/satori/go.uuid"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithID(id *uuid.UUID) BlockBuilder
	WithBlock(blk blocks.SignedBlock) BlockBuilder
	WithSignatures(sigs []users.Signature) BlockBuilder
	CreatedOn(ts time.Time) BlockBuilder
	Now() (Block, error)
}
