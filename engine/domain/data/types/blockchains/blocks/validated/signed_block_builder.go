package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	users "github.com/XMNBlockchain/openxmn/engine/domain/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// SignedBlockBuilder represents a signed block builder
type SignedBlockBuilder interface {
	Create() SignedBlockBuilder
	WithID(id *uuid.UUID) SignedBlockBuilder
	WithMetaData(met metadata.MetaData) SignedBlockBuilder
	WithBlock(blk Block) SignedBlockBuilder
	WithSignature(sig users.Signature) SignedBlockBuilder
	CreatedOn(ts time.Time) SignedBlockBuilder
	Now() (SignedBlock, error)
}
