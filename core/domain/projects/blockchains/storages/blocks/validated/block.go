package validated

import (
	stored_block "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
	stored_users "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/users"
)

// Block represents a stored validated block
type Block interface {
	GetMetaData() stored_files.File
	GetBlock() stored_block.SignedBlock
	GetSignatures() stored_users.Signatures
}
