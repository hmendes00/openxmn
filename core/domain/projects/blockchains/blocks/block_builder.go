package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/metadata"
	aggregated "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/transactions/signed/aggregated"
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
