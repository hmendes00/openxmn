package domain

import (
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
)

// TransactionsBuilder represents the Transactions builder
type TransactionsBuilder interface {
	Create() TransactionsBuilder
	WithTransactions(trs []signed_transactions.Transaction) TransactionsBuilder
	WithAtomicTransactions(trs []signed_transactions.AtomicTransaction) TransactionsBuilder
	Now() (Transactions, error)
}
