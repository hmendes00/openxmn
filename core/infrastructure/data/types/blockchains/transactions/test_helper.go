package transactions

import (
	"encoding/json"
	"strconv"
	"time"

	trs "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/transactions"
	concrete_chunks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/chunks"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/hashtrees"
	concrete_met "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	concrete_stored_trs "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/blockchains/transactions"
	uuid "github.com/satori/go.uuid"
)

// JsDataForTests represents a structure for tests
type JsDataForTests struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateTransactionForTests creates a Transaction for tests
func CreateTransactionForTests() *Transaction {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		js,
	}
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_met.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	trs := createTransaction(met.(*concrete_met.MetaData), js)
	return trs.(*Transaction)
}

// CreateTransactionsForTests creates a Transactions for tests
func CreateTransactionsForTests() *Transactions {
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := []*Transaction{
		CreateTransactionForTests(),
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
	met, _ := concrete_met.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	out := createTransactions(met.(*concrete_met.MetaData), trs)
	return out.(*Transactions)
}

// CreateTransactionBuilderFactoryForTests creates a new TransactionBuilderFactory for tests
func CreateTransactionBuilderFactoryForTests() trs.TransactionBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metBuilderFactory := concrete_met.CreateMetaDataBuilderFactoryForTests()
	out := CreateTransactionBuilderFactory(htBuilderFactory, metBuilderFactory)
	return out
}

// CreateTransactionRepositoryForTests creates a new TransactionRepository for tests
func CreateTransactionRepositoryForTests() trs.TransactionRepository {
	chksRepository := concrete_chunks.CreateRepositoryForTests()
	metRepository := concrete_met.CreateMetaDataRepositoryForTests()
	transactionBuilderFactory := CreateTransactionBuilderFactoryForTests()
	out := CreateTransactionRepository(chksRepository, metRepository, transactionBuilderFactory)
	return out
}

// CreateTransactionServiceForTests creates a new TransactionService for tests
func CreateTransactionServiceForTests() trs.TransactionService {
	metaDataService := concrete_met.CreateMetaDataServiceForTests()
	chkBuilderFactory := concrete_chunks.CreateBuilderFactoryForTests()
	chkService := concrete_chunks.CreateServiceForTests()
	storedTrsBuilderFactory := concrete_stored_trs.CreateTransactionBuilderFactoryForTests()
	out := CreateTransactionService(metaDataService, chkBuilderFactory, chkService, storedTrsBuilderFactory)
	return out
}

// CreateTransactionsBuilderFactoryForTests creates a new TransactionsBuilderFactory for tests
func CreateTransactionsBuilderFactoryForTests() trs.TransactionsBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metBuilderFactory := concrete_met.CreateMetaDataBuilderFactoryForTests()
	out := CreateTransactionsBuilderFactory(htBuilderFactory, metBuilderFactory)
	return out
}

// CreateTransactionsRepositoryForTests creates a new TransactionsRepository for tests
func CreateTransactionsRepositoryForTests() trs.TransactionsRepository {
	metRepository := concrete_met.CreateMetaDataRepositoryForTests()
	trsRepository := CreateTransactionRepositoryForTests()
	trsBuilderFactory := CreateTransactionsBuilderFactoryForTests()
	out := CreateTransactionsRepository(metRepository, trsRepository, trsBuilderFactory)
	return out
}

// CreateTransactionsServiceForTests creates a new TransactionsService for tests
func CreateTransactionsServiceForTests() trs.TransactionsService {
	metaDataService := concrete_met.CreateMetaDataServiceForTests()
	trsService := CreateTransactionServiceForTests()
	storedTrsBuilderFactory := concrete_stored_trs.CreateBuilderFactory()
	out := CreateTransactionsService(metaDataService, trsService, storedTrsBuilderFactory)
	return out
}
