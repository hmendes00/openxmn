package domain

import (
	transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/domain"
	users "github.com/XMNBlockchain/core/packages/users/domain"
	uuid "github.com/satori/go.uuid"
)

// AtomicTransactionBuilder represents an AtomicTransaction builder
type AtomicTransactionBuilder interface {
	Create() AtomicTransactionBuilder
	WithID(id *uuid.UUID) AtomicTransactionBuilder
	WithTransactions(trs []transactions.Transaction) AtomicTransactionBuilder
	WithSignature(sig users.Signature) AtomicTransactionBuilder
	Now() (AtomicTransaction, error)
}