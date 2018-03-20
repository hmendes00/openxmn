package signed

import (
	"strconv"
	"time"

	signed_transactions "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/transactions/signed"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/transactions/signed"
	concrete_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests() *Transaction {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetHashTree().GetHash().Get(),
	}
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	sigTrs := createTransaction(met.(*concrete_metadata.MetaData), trs, sig)
	return sigTrs.(*Transaction)
}

// CreateTransactionsForTests creates a Transactions for tests
func CreateTransactionsForTests() *Transactions {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := []*Transaction{
		CreateTransactionForTests(),
		CreateTransactionForTests(),
		CreateTransactionForTests(),
		CreateTransactionForTests(),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}

	for _, oneTrs := range trs {
		blocks = append(blocks, oneTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	sigTrs := createTransactions(met.(*concrete_metadata.MetaData), trs)
	return sigTrs.(*Transactions)
}

// CreateAtomicTransactionForTests creates an AtomicTransaction for tests
func CreateAtomicTransactionForTests() *AtomicTransaction {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	atomicTrs := createAtomicTransaction(met.(*concrete_metadata.MetaData), trs, sig)
	return atomicTrs.(*AtomicTransaction)
}

// CreateAtomicTransactionsForTests creates an AtomicTransaction for tests
func CreateAtomicTransactionsForTests() *AtomicTransactions {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	atomicTrs := []*AtomicTransaction{
		CreateAtomicTransactionForTests(),
		CreateAtomicTransactionForTests(),
		CreateAtomicTransactionForTests(),
		CreateAtomicTransactionForTests(),
		CreateAtomicTransactionForTests(),
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}

	for _, oneAtomicTrs := range atomicTrs {
		blocks = append(blocks, oneAtomicTrs.GetMetaData().GetHashTree().GetHash().Get())
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	out := createAtomicTransactions(met.(*concrete_metadata.MetaData), atomicTrs)
	return out.(*AtomicTransactions)
}

// CreateTransactionBuilderFactoryForTests creates a new TransactionBuilderFactory for tests
func CreateTransactionBuilderFactoryForTests() signed_transactions.TransactionBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateTransactionRepositoryForTests creates a new TransactionRepository for tests
func CreateTransactionRepositoryForTests() signed_transactions.TransactionRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	sigRepository := concrete_users.CreateSignatureRepositoryForTests()
	transactionRepository := concrete_transactions.CreateTransactionRepositoryForTests()
	signedTrsBuilderFactory := CreateTransactionBuilderFactoryForTests()
	out := CreateTransactionRepository(metaDataRepository, sigRepository, transactionRepository, signedTrsBuilderFactory)
	return out
}

// CreateTransactionServiceForTests creates a new TransactionService for tests
func CreateTransactionServiceForTests() signed_transactions.TransactionService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	trsService := concrete_transactions.CreateTransactionServiceForTests()
	storedSignedTrsBuilderFactory := concrete_stored_signed_transactions.CreateTransactionBuilderFactoryForTests()
	sigService := concrete_users.CreateSignatureServiceForTests()
	out := CreateTransactionService(metaDataService, trsService, storedSignedTrsBuilderFactory, sigService)
	return out
}

// CreateTransactionsBuilderFactoryForTests creates a new TransactionsBuilderFactory for tests
func CreateTransactionsBuilderFactoryForTests() signed_transactions.TransactionsBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateTransactionsRepositoryForTests creates a new TransactionsRepository for tests
func CreateTransactionsRepositoryForTests() signed_transactions.TransactionsRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	trsRepository := CreateTransactionRepositoryForTests()
	signedTrsBuilderFactory := CreateTransactionsBuilderFactoryForTests()
	out := CreateTransactionsRepository(metaDataRepository, trsRepository, signedTrsBuilderFactory)
	return out
}

// CreateTransactionsServiceForTests creates a new TransactionsService for tests
func CreateTransactionsServiceForTests() signed_transactions.TransactionsService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	trsService := CreateTransactionServiceForTests()
	storedTrsBuilderFactory := concrete_stored_signed_transactions.CreateTransactionsBuilderFactoryForTests()
	out := CreateTransactionsService(metaDataService, trsService, storedTrsBuilderFactory)
	return out
}

// CreateAtomicTransactionBuilderFactoryForTests creates a new AtomicTransactionBuilderFactory for tests
func CreateAtomicTransactionBuilderFactoryForTests() signed_transactions.AtomicTransactionBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateAtomicTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateAtomicTransactionRepositoryForTests creates a new AtomicTransactionRepository for tests
func CreateAtomicTransactionRepositoryForTests() signed_transactions.AtomicTransactionRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	userSigRepository := concrete_users.CreateSignatureRepositoryForTests()
	trsRepository := concrete_transactions.CreateTransactionsRepositoryForTests()
	atomicTrsBuilderFactory := CreateAtomicTransactionBuilderFactoryForTests()
	out := CreateAtomicTransactionRepository(metaDataRepository, userSigRepository, trsRepository, atomicTrsBuilderFactory)
	return out
}

// CreateAtomicTransactionServiceForTests creates a new AtomicTransactionService for tests
func CreateAtomicTransactionServiceForTests() signed_transactions.AtomicTransactionService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	trsService := concrete_transactions.CreateTransactionsServiceForTests()
	storedAtomicTrsBuilderFactory := concrete_stored_signed_transactions.CreateAtomicTransactionBuilderFactoryForTests()
	userSigService := concrete_users.CreateSignatureServiceForTests()
	out := CreateAtomicTransactionService(metaDataService, userSigService, trsService, storedAtomicTrsBuilderFactory)
	return out
}

// CreateAtomicTransactionsBuilderFactoryForTests creates a new AtomicTransactionsBuilderFactory for tests
func CreateAtomicTransactionsBuilderFactoryForTests() signed_transactions.AtomicTransactionsBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateAtomicTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateAtomicTransactionsRepositoryForTests creates a new AtomicTransactionsRepository for tests
func CreateAtomicTransactionsRepositoryForTests() signed_transactions.AtomicTransactionsRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	atomicTrsRepository := CreateAtomicTransactionRepositoryForTests()
	atomicTransBuilderFactory := CreateAtomicTransactionsBuilderFactoryForTests()
	out := CreateAtomicTransactionsRepository(metaDataRepository, atomicTrsRepository, atomicTransBuilderFactory)
	return out
}

// CreateAtomicTransactionsServiceForTests creates a new AtomicTransactionsService for tests
func CreateAtomicTransactionsServiceForTests() signed_transactions.AtomicTransactionsService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	atomicTrsService := CreateAtomicTransactionServiceForTests()
	storedAtomicTrsBuilderFactory := concrete_stored_signed_transactions.CreateAtomicTransactionsBuilderFactoryForTests()
	out := CreateAtomicTransactionsService(metaDataService, atomicTrsService, storedAtomicTrsBuilderFactory)
	return out
}
