package domain

import (
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// TransactionBuilder represents a Transaction builder
type TransactionBuilder interface {
	Create() TransactionBuilder
	WithMetaData(met stored_files.File) TransactionBuilder
	WithChunks(chks stored_chunks.Chunks) TransactionBuilder
	Now() (Transaction, error)
}
