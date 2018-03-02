package domain

import (
	validated "github.com/XMNBlockchain/core/packages/blockchains/blocks/validated/domain"
	stored_chains "github.com/XMNBlockchain/core/packages/storages/chains/domain"
)

// ChainService represents a chain service
type ChainService interface {
	Save(ch Chain) (stored_chains.Chain, error)
	Chain(validatedBlk validated.Block) (Chain, error)
}
