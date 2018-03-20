package transactions

import (
	stored_chunks "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/chunks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
)

// Transaction represents a stored transaction
type Transaction interface {
	GetMetaData() stored_files.File
	GetChunks() stored_chunks.Chunks
}
