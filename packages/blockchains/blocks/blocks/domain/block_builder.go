package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/domain"
	uuid "github.com/satori/go.uuid"
)

// BlockBuilder represents a block builder
type BlockBuilder interface {
	Create() BlockBuilder
	WithID(id *uuid.UUID) BlockBuilder
	WithMetaData(met metadata.MetaData) BlockBuilder
	WithTransactions(trs []aggregated.SignedTransactions) BlockBuilder
	CreatedOn(crOn time.Time) BlockBuilder
	Now() (Block, error)
}
