package domain

import (
	stored_chained_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/chained/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// Chain represents a stored chain
type Chain interface {
	GetMetaData() stored_files.File
	GetHashTree() stored_files.File
	GetFloorBlock() stored_chained_blocks.Block
	GetCeilBlock() stored_chained_blocks.Block
}
