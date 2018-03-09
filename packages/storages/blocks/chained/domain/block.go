package domain

import (
	stored_validated_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/validated/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Block represents a stored chained block
type Block interface {
	GetMetaData() stored_files.File
	GetBlock() stored_validated_blocks.Block
}
