package signed

import (
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed"
)

// CreateTransactionBuilderFactoryForTests creates a TransactionBuilderFactory for tests
func CreateTransactionBuilderFactoryForTests() stored_signed_transactions.TransactionBuilderFactory {
	out := CreateTransactionBuilderFactory()
	return out
}

// CreateTransactionsBuilderFactoryForTests creates a TransactionsBuilderFactory for tests
func CreateTransactionsBuilderFactoryForTests() stored_signed_transactions.TransactionsBuilderFactory {
	out := CreateTransactionsBuilderFactory()
	return out
}

// CreateAtomicTransactionBuilderFactoryForTests creates a AtomicTransactionBuilderFactory for tests
func CreateAtomicTransactionBuilderFactoryForTests() stored_signed_transactions.AtomicTransactionBuilderFactory {
	out := CreateAtomicTransactionBuilderFactory()
	return out
}

// CreateAtomicTransactionsBuilderFactoryForTests creates a AtomicTransactionsBuilderFactory for tests
func CreateAtomicTransactionsBuilderFactoryForTests() stored_signed_transactions.AtomicTransactionsBuilderFactory {
	out := CreateAtomicTransactionsBuilderFactory()
	return out
}
