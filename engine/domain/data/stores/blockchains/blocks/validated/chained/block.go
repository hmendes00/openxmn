package chained

import (
	stored_validated_blocks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/blocks/validated"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Block represents a stored chained block
type Block interface {
	GetMetaData() stored_files.File
	GetBlock() stored_validated_blocks.Block
}
