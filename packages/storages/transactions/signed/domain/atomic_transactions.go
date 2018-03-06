package domain

import (
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// AtomicTransactions represents stored signed AtomicTransactions
type AtomicTransactions interface {
	GetMetaData() stored_files.File
	GetTransactions() []AtomicTransaction
}
