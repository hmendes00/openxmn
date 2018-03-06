package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/domain"
	uuid "github.com/satori/go.uuid"
)

// AggregatedSignedTransactionsBuilder represents an aggregated signed transactions builder
type AggregatedSignedTransactionsBuilder interface {
	Create() AggregatedSignedTransactionsBuilder
	WithID(id *uuid.UUID) AggregatedSignedTransactionsBuilder
	WithMetaData(met metadata.MetaData) AggregatedSignedTransactionsBuilder
	WithTransactions(trs []Transactions) AggregatedSignedTransactionsBuilder
	CreatedOn(crOn time.Time) AggregatedSignedTransactionsBuilder
	Now() (AggregatedSignedTransactions, error)
}
