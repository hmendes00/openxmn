package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// SignedTransactionsService represents a signed transactions service
type SignedTransactionsService interface {
	Save(dirPath string, trs SignedTransactions) (stored_objects.Tree, error)
	SaveAll(dirPath string, trs []SignedTransactions) ([]stored_objects.Tree, error)
}
