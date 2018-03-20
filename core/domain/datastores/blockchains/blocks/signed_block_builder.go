package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

// SignedBlockBuilder represents a SignedBlock builder
type SignedBlockBuilder interface {
	Create() SignedBlockBuilder
	WithID(id *uuid.UUID) SignedBlockBuilder
	WithMetaData(met metadata.MetaData) SignedBlockBuilder
	WithBlock(blk Block) SignedBlockBuilder
	WithSignature(sig users.Signature) SignedBlockBuilder
	CreatedOn(ts time.Time) SignedBlockBuilder
	Now() (SignedBlock, error)
}
