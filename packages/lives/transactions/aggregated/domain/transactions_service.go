package domain

import (
	stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/domain"
)

// TransactionsService represents a transactions service
type TransactionsService interface {
	Save(dirPath string, trs Transactions) (stored_objects.Tree, error)
	SaveAll(dirPath string, trs []Transactions) ([]stored_objects.Tree, error)
}
