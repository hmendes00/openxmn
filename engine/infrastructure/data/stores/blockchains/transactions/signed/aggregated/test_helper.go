package aggregated

import (
	stored_aggregated_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions/signed/aggregated"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/transactions/signed"
	concrete_stored_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/users"
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

// CreateTransactionsRepositoryForTests creates a TransactionsRepository for tests
func CreateTransactionsRepositoryForTests() stored_aggregated_transactions.TransactionsRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	transRepository := concrete_stored_signed_transactions.CreateTransactionsRepositoryForTests()
	atomicTransRepository := concrete_stored_signed_transactions.CreateAtomicTransactionsRepositoryForTests()
	aggrTransBuilderFactory := CreateTransactionsBuilderFactoryForTests()
	out := CreateTransactionsRepository(fileRepository, transRepository, atomicTransRepository, aggrTransBuilderFactory)
	return out
}

// CreateSignedTransactionsRepositoryForTests creates a SignedTransactionsRepository for tests
func CreateSignedTransactionsRepositoryForTests() stored_aggregated_transactions.SignedTransactionsRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	sigRepository := concrete_stored_users.CreateSignatureRepositoryForTests()
	transRepository := CreateTransactionsRepositoryForTests()
	signedTransBuilderFactory := CreateSignedTransactionsBuilderFactoryForTests()
	out := CreateSignedTransactionsRepository(fileRepository, sigRepository, transRepository, signedTransBuilderFactory)
	return out
}
