package blocks

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/users"
)

// SignedBlockBuilder represents a stored signed block builder
type SignedBlockBuilder interface {
	Create() SignedBlockBuilder
	WithMetaData(met stored_files.File) SignedBlockBuilder
	WithSignature(sig stored_users.Signature) SignedBlockBuilder
	WithBlock(blk Block) SignedBlockBuilder
	Now() (SignedBlock, error)
}