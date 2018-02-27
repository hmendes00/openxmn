package infrastructure

import (
	trs "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
)

type transactionBuilderFactory struct {
}

// CreateTransactionBuilderFactory creates a new TransactionBuilderFactory
func CreateTransactionBuilderFactory() trs.TransactionBuilderFactory {
	out := transactionBuilderFactory{}
	return &out
}

// Create creates a new TransactionBuilder instance
func (fac *transactionBuilderFactory) Create() trs.TransactionBuilder {
	out := createTransactionBuilder()
	return out
}
