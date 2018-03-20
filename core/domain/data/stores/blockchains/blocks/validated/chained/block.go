package chained

import (
	stored_validated_blocks "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/blockchains/blocks/validated"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
)

// Block represents a stored chained block
type Block interface {
	GetMetaData() stored_files.File
	GetBlock() stored_validated_blocks.Block
}
