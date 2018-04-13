package validated

import (
	stored_users "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/users"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// SignedBlockBuilder represents a stored validated SignedBlockBuilder
type SignedBlockBuilder interface {
	Create() SignedBlockBuilder
	WithMetaData(met stored_files.File) SignedBlockBuilder
	WithBlock(blk Block) SignedBlockBuilder
	WithSignature(sig stored_users.Signature) SignedBlockBuilder
	Now() (SignedBlock, error)
}
