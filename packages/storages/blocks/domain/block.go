package domain

import (
	"time"

	hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
)

// Block represents a stored block
type Block interface {
	GetHashTree() hashtrees.HashTree
	GetSignature() users.Signature
	GetTransactions() []stored_aggregated_transactions.Transactions
	GetNeededKarma() int
	CreatedOn() time.Time
}
