package domain

import (
	"time"

	blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/domain"
	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	uuid "github.com/satori/go.uuid"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithID(id *uuid.UUID) BlockBuilder
	WithMetaData(met metadata.MetaData) BlockBuilder
	WithBlock(blk blocks.SignedBlock) BlockBuilder
	WithSignatures(sigs users.Signatures) BlockBuilder
	CreatedOn(ts time.Time) BlockBuilder
	Now() (Block, error)
}
