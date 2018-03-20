package signed

import (
	stored_files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
)

// AtomicTransactions represents stored signed AtomicTransactions
type AtomicTransactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []AtomicTransaction
}
