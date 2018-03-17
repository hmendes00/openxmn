package aggregated

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/transactions/signed/aggregated"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/transactions/signed"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/users"
)

// CreateTransactionsForTests creates a Transactions for tests
func CreateTransactionsForTests() *Transactions {
	met := concrete_stored_files.CreateFileForTests()
	trs := concrete_stored_signed_transactions.CreateTransactionsForTests()
	atomicTrs := concrete_stored_signed_transactions.CreateAtomicTransactionsForTests()
	out := createTransactions(met, trs, atomicTrs)
	return out.(*Transactions)
}

// CreateSignedTransactionsForTests creates a SignedTransactions for tests
func CreateSignedTransactionsForTests() *SignedTransactions {
	met := concrete_stored_files.CreateFileForTests()
	sig := concrete_stored_users.CreateSignatureForTests()
	trs := CreateTransactionsForTests()
	out := createSignedTransactions(met, sig, trs)
	return out.(*SignedTransactions)
}

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
