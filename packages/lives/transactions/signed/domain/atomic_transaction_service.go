package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// AtomicTransactionService represents a transaction service
type AtomicTransactionService interface {
	Save(dirPath string, trs AtomicTransaction) (stored_objects.Tree, error)
	SaveAll(dirPath string, trs []AtomicTransaction) ([]stored_objects.Tree, error)
}
