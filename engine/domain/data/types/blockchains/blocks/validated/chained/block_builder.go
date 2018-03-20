package domain

import (
	"time"

	validated "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/blocks/validated"
	uuid "github.com/satori/go.uuid"
)

// BlockBuilder represents a chained block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithID(id *uuid.UUID) BlockBuilder
	WithMetaData(met MetaData) BlockBuilder
	WithBlock(blk validated.Block) BlockBuilder
	WithPreviousID(prevID *uuid.UUID) BlockBuilder
	CreatedOn(crOn time.Time) BlockBuilder
	Now() (Block, error)
}
