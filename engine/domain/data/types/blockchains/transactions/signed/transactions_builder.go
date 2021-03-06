package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	uuid "github.com/satori/go.uuid"
)

// TransactionsBuilder represents a TransactionsBuilder instance
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithID(id *uuid.UUID) TransactionsBuilder
	WithMetaData(met metadata.MetaData) TransactionsBuilder
	WithTransactions(trs []Transaction) TransactionsBuilder
	CreatedOn(crOn time.Time) TransactionsBuilder
	Now() (Transactions, error)
}
