package domain

import (
	signed_transactions "github.com/XMNBlockchain/core/packages/transactions/signed/domain"
)

// Transactions represents aggregated signed transactions
type Transactions interface {
	GetTrs() []signed_transactions.Transaction
	GetAtomicTrs() []signed_transactions.AtomicTransaction
}
