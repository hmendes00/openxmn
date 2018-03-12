package aggregated

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
)

// CreateTransactionsBuilderFactoryForTests creates a TransactionsBuilderFactory for tests
func CreateTransactionsBuilderFactoryForTests() stored_aggregated_transactions.TransactionsBuilderFactory {
	out := CreateTransactionsBuilderFactory()
	return out
}

// CreateSignedTransactionsBuilderFactoryForTests creates a SignedTransactionsBuilderFactory for tests
func CreateSignedTransactionsBuilderFactoryForTests() stored_aggregated_transactions.SignedTransactionsBuilderFactory {
	out := CreateSignedTransactionsBuilderFactory()
	return out
}
