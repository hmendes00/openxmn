package domain

import (
	aggregated "github.com/XMNBlockchain/core/packages/transactions/aggregated/domain"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithTransactions(trs []aggregated.SignedTransactions) BlockBuilder
	Now() (Block, error)
}
