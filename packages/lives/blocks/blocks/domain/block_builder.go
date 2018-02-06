package domain

import (
	aggregated "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/domain"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithTransactions(trs []aggregated.SignedTransactions) BlockBuilder
	Now() (Block, error)
}
