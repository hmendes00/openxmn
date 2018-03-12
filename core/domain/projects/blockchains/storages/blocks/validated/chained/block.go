package domain

import (
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/blocks/validated"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// Block represents a stored chained block
type Block interface {
	GetMetaData() stored_files.File
	GetBlock() stored_validated_blocks.Block
}
