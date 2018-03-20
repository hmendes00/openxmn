package validated

import (
	stored_block "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/users"
)

// BlockBuilder represents a stored validated block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithMetaData(met stored_files.File) BlockBuilder
	WithBlock(blk stored_block.SignedBlock) BlockBuilder
	WithSignatures(sigs stored_users.Signatures) BlockBuilder
	Now() (Block, error)
}
