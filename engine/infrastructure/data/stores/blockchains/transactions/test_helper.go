package transactions

import (
	stored_transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/blockchains/transactions"
	concrete_stored_chunks "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/chunks"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests() *Transaction {
	met := concrete_stored_files.CreateFileForTests()
	chks := concrete_stored_chunks.CreateChunksForTests()
	out := createTransaction(met, chks)
	return out.(*Transaction)
}

// CreateTransactionsForTests creates a Transactions for tests
func CreateTransactionsForTests() *Transactions {
	met := concrete_stored_files.CreateFileForTests()
	trs := []*Transaction{
		CreateTransactionForTests(),
		CreateTransactionForTests(),
		CreateTransactionForTests(),
		CreateTransactionForTests(),
	}

	out := createTransactions(met, trs)
	return out.(*Transactions)
}

// CreateTransactionBuilderFactoryForTests creates a new TransactionBuilderFactory for tests
func CreateTransactionBuilderFactoryForTests() stored_transactions.TransactionBuilderFactory {
	out := CreateTransactionBuilderFactory()
	return out
}

// CreateBuilderFactoryForTests creates a new TransactionsBuilderFactory for tests
func CreateBuilderFactoryForTests() stored_transactions.BuilderFactory {
	out := CreateBuilderFactory()
	return out
}

// CreateTransactionRepositoryForTests creates a TransactionRepository for tests
func CreateTransactionRepositoryForTests() stored_transactions.TransactionRepository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	chkRepository := concrete_stored_chunks.CreateRepositoryForTests()
	trsBuilderFactory := CreateTransactionBuilderFactoryForTests()
	out := CreateTransactionRepository(fileRepository, chkRepository, trsBuilderFactory)
	return out
}

// CreateRepositoryForTests creates a Repository for tests
func CreateRepositoryForTests() stored_transactions.Repository {
	fileRepository := concrete_stored_files.CreateFileRepositoryForTests()
	trsRepository := CreateTransactionRepositoryForTests()
	transBuilderFactory := CreateBuilderFactoryForTests()
	out := CreateRepository(fileRepository, trsRepository, transBuilderFactory)
	return out
}
