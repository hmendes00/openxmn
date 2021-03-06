package domain

import (
	"time"

	blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks"
	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
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
