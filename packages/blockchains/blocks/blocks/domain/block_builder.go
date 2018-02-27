package domain

import (
	"time"

	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	uuid "github.com/satori/go.uuid"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithID(id *uuid.UUID) BlockBuilder
	WithTransactions(trs []aggregated.SignedTransactions) BlockBuilder
	CreatedOn(ts time.Time) BlockBuilder
	Now() (Block, error)
}
