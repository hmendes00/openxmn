package transactions

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
)

// Builder represents a Transactions builder
type Builder interface {
	Create() Builder
	WithMetaData(met stored_files.File) Builder
	WithTransactions(trs []Transaction) Builder
	Now() (Transactions, error)
}
