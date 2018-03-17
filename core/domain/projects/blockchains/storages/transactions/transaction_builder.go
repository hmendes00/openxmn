package transactions

import (
	stored_chunks "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/chunks"
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// TransactionBuilder represents a Transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithMetaData(met stored_files.File) TransactionBuilder
	WithChunks(chks stored_chunks.Chunks) TransactionBuilder
	Now() (Transaction, error)
}
