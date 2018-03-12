package transactions

import (
	stored_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions"
)

// CreateTransactionBuilderFactoryForTests creates a new TransactionBuilderFactory for tests
func CreateTransactionBuilderFactoryForTests() stored_transactions.TransactionBuilderFactory {
	out := CreateTransactionBuilderFactory()
	return out
}

// CreateTransactionsBuilderFactoryForTests creates a new TransactionsBuilderFactory for tests
func CreateTransactionsBuilderFactoryForTests() stored_transactions.TransactionsBuilderFactory {
	out := CreateTransactionsBuilderFactory()
	return out
}
