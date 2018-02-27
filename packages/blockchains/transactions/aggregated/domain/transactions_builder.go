package domain

import (
	"time"

	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	uuid "github.com/satori/go.uuid"
)

// TransactionsBuilder represents the Transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithID(id *uuid.UUID) TransactionsBuilder
	WithTransactions(trs []signed_transactions.Transaction) TransactionsBuilder
	WithAtomicTransactions(trs []signed_transactions.AtomicTransaction) TransactionsBuilder
	CreatedOn(ts time.Time) TransactionsBuilder
	Now() (Transactions, error)
}
