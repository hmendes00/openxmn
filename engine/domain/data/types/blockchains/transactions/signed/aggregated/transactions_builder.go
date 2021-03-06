package domain

import (
	"time"

	metadata "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	signed_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed"
	uuid "github.com/satori/go.uuid"
)

// TransactionsBuilder represents the Transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithID(id *uuid.UUID) TransactionsBuilder
	WithMetaData(met metadata.MetaData) TransactionsBuilder
	WithTransactions(trs signed_transactions.Transactions) TransactionsBuilder
	WithAtomicTransactions(trs signed_transactions.AtomicTransactions) TransactionsBuilder
	CreatedOn(ts time.Time) TransactionsBuilder
	Now() (Transactions, error)
}
