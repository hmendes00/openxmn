package signed

import (
	stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/transactions/signed"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
	concrete_stored_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/transactions"
	concrete_stored_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/users"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests() *Transaction {
	met := concrete_stored_files.CreateFileForTests()
	sig := concrete_stored_users.CreateSignatureForTests()
	trs := concrete_stored_transactions.CreateTransactionForTests()
	out := createTransaction(met, sig, trs)
	return out.(*Transaction)
}

// CreateTransactionsForTests creates a Transactions for tests
func CreateTransactionsForTests() *Transactions {
	met := concrete_stored_files.CreateFileForTests()
	trs := []*Transaction{
		CreateTransactionForTests(),
		CreateTransactionForTests(),
		CreateTransactionForTests(),
	}
	out := createTransactions(met, trs)
	return out.(*Transactions)
}

// CreateAtomicTransactionForTests creates a AtomicTransaction for tests
func CreateAtomicTransactionForTests() *AtomicTransaction {
	met := concrete_stored_files.CreateFileForTests()
	sig := concrete_stored_users.CreateSignatureForTests()
	trs := concrete_stored_transactions.CreateTransactionsForTests()
	out := createAtomicTransaction(met, sig, trs)
	return out.(*AtomicTransaction)
}

// CreateAtomicTransactionsForTests creates a AtomicTransactions for tests
func CreateAtomicTransactionsForTests() *AtomicTransactions {
	met := concrete_stored_files.CreateFileForTests()
	trs := []*AtomicTransaction{
		CreateAtomicTransactionForTests(),
		CreateAtomicTransactionForTests(),
		CreateAtomicTransactionForTests(),
	}
	out := createAtomicTransactions(met, trs)
	return out.(*AtomicTransactions)
}

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

// CreateTransactionRepositoryForTests creates a TransactionRepository for tests
func CreateTransactionRepositoryForTests() stored_signed_transactions.TransactionRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	sigRepository := concrete_stored_users.CreateSignatureRepositoryForTests()
	trsRepository := concrete_stored_transactions.CreateTransactionRepositoryForTests()
	signedTrsBuilderFactory := CreateTransactionBuilderFactoryForTests()
	out := CreateTransactionRepository(fileRepository, sigRepository, trsRepository, signedTrsBuilderFactory)
	return out
}

// CreateTransactionsRepositoryForTests creates a TransactionsRepository for tests
func CreateTransactionsRepositoryForTests() stored_signed_transactions.TransactionsRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	signedTrsRepository := CreateTransactionRepositoryForTests()
	signedTransBuilderFactory := CreateTransactionsBuilderFactoryForTests()
	out := CreateTransactionsRepository(fileRepository, signedTrsRepository, signedTransBuilderFactory)
	return out
}

// CreateAtomicTransactionRepositoryForTests creates a AtomicTransactionRepository for tests
func CreateAtomicTransactionRepositoryForTests() stored_signed_transactions.AtomicTransactionRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	sigRepository := concrete_stored_users.CreateSignatureRepositoryForTests()
	transRepository := concrete_stored_transactions.CreateRepositoryForTests()
	atomicTrsBuilderFactory := CreateAtomicTransactionBuilderFactoryForTests()
	out := CreateAtomicTransactionRepository(fileRepository, sigRepository, transRepository, atomicTrsBuilderFactory)
	return out
}

// CreateAtomicTransactionsRepositoryForTests creates a AtomicTransactionsRepository for tests
func CreateAtomicTransactionsRepositoryForTests() stored_signed_transactions.AtomicTransactionsRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	atomicTrsRepository := CreateAtomicTransactionRepositoryForTests()
	atomicTransBuilderFactory := CreateAtomicTransactionsBuilderFactoryForTests()
	out := CreateAtomicTransactionsRepository(fileRepository, atomicTrsRepository, atomicTransBuilderFactory)
	return out
}
