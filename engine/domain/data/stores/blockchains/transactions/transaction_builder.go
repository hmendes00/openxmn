package transactions

import (
	stored_chunks "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/chunks"
	stored_files "github.com/XMNBlockchain/exmachina-network/engine/domain/data/stores/files"
)

// TransactionBuilder represents a Transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithMetaData(met stored_files.File) TransactionBuilder
	WithChunks(chks stored_chunks.Chunks) TransactionBuilder
	Now() (Transaction, error)
}
