package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// SignedBlockBuilder represents a stored signed block builder
type SignedBlockBuilder interface {
	Create() SignedBlockBuilder
	WithMetaData(met stored_files.File) SignedBlockBuilder
	WithSignature(sig stored_files.File) SignedBlockBuilder
	WithBlock(blk Block) SignedBlockBuilder
	Now() (SignedBlock, error)
}