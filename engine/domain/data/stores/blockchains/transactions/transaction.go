package transactions

import (
	stored_chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// Transaction represents a stored transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetChunks() stored_chunks.Chunks
}
