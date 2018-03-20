package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// AtomicTransactionsBuilder represents an AtomicTransactions builder
type AtomicTransactionsBuilder interface {
	Create() AtomicTransactionsBuilder
	WithID(id *uuid.UUID) AtomicTransactionsBuilder
	WithMetaData(met metadata.MetaData) AtomicTransactionsBuilder
	WithTransactions(trs []AtomicTransaction) AtomicTransactionsBuilder
	CreatedOn(crOn time.Time) AtomicTransactionsBuilder
	Now() (AtomicTransactions, error)
}
