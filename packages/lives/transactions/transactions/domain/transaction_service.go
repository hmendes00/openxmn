package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TransactionService represents a transaction service
type TransactionService interface {
	Save(dirPath string, trs Transaction) (stored_objects.Object, error)
	SaveAll(dirPath string, trs []Transaction) ([]stored_objects.Object, error)
}
