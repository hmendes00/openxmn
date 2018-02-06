package domain

import (
	hashtree "github.com/XMNBlockchain/core/packages/hashtrees/domain"
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
)

// Block represents a Block instance
type Block interface {
	GetHashTree() hashtree.HashTree
	GetTransactions() []aggregated.SignedTransactions
	GetNeededKarma() int
}
