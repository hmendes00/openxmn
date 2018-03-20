package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// TransactionsBuilder represents a Transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithID(id *uuid.UUID) TransactionsBuilder
	WithMetaData(meta metadata.MetaData) TransactionsBuilder
	WithTransactions(trs []Transaction) TransactionsBuilder
	CreatedOn(crOn time.Time) TransactionsBuilder
	Now() (Transactions, error)
}
