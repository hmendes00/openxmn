package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_chunks "github.com/XMNBlockchain/core/packages/blockchains/chunks/infrastructure"
	concrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	concrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/infrastructure"
	concrete_stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/infrastructure"
)

func TestSaveAtomicTransactions_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateAtomicTransactionsForTests(t)

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 16
	extension := "chk"

	//factories:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	storedChkBuilderFactory := concrete_stored_chunks.CreateChunksBuilderFactory()
	fileRepository := concrete_files.CreateFileRepository(fileBuilderFactory)
	fileService := concrete_files.CreateFileService(storedFileBuilderFactory)
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	htService := concrete_hashtrees.CreateHashTreeService(fileService, fileBuilderFactory)
	htRepository := concrete_hashtrees.CreateHashTreeRepository(fileRepository)
	chksBuilderFactory := concrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := concrete_chunks.CreateChunksRepository(htRepository, fileRepository, chksBuilderFactory)
	chkService := concrete_chunks.CreateChunksService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	metaDataRepository := concrete_metadata.CreateMetaDataRepository(fileRepository)
	metaDataService := concrete_metadata.CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	trsBuilderFactory := concrete_transactions.CreateTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	transBuilderFactory := concrete_transactions.CreateTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	trsRepository := concrete_transactions.CreateTransactionRepository(chkRepository, metaDataRepository, trsBuilderFactory)
	storedTrsBuilderFactory := concrete_stored_transactions.CreateTransactionBuilderFactory()
	storedTrsansBuilderFactory := concrete_stored_transactions.CreateTransactionsBuilderFactory()
	trsService := concrete_transactions.CreateTransactionService(metaDataService, chksBuilderFactory, chkService, storedTrsBuilderFactory)
	userSigRepository := concrete_users.CreateSignatureRepository(fileRepository)
	userSigService := concrete_users.CreateSignatureService(fileService, fileBuilderFactory)
	transRepository := concrete_transactions.CreateTransactionsRepository(metaDataRepository, trsRepository, transBuilderFactory)
	transService := concrete_transactions.CreateTransactionsService(metaDataService, trsService, storedTrsansBuilderFactory)
	atomicTrsBuilderFactory := CreateAtomicTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedAtomicTrsBuilderFactory := concrete_stored_signed_transactions.CreateAtomicTransactionBuilderFactory()
	atomicTrsRepository := CreateAtomicTransactionRepository(metaDataRepository, userSigRepository, transRepository, atomicTrsBuilderFactory)
	atomicTrsService := CreateAtomicTransactionService(metaDataService, userSigService, transService, storedAtomicTrsBuilderFactory)
	atomicTransBuilderFactory := CreateAtomicTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedAtomicTransBuilderFactory := concrete_stored_signed_transactions.CreateAtomicTransactionsBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateAtomicTransactionsRepository(metaDataRepository, atomicTrsRepository, atomicTransBuilderFactory)
	service := CreateAtomicTransactionsService(metaDataService, atomicTrsService, storedAtomicTransBuilderFactory)

	//make sure there is no transaction:
	_, noTrsErr := repository.Retrieve(basePath)
	if noTrsErr == nil {
		t.Errorf("there was supposed to be no transaction.")
	}

	//save the transaction:
	_, storedTrsErr := service.Save(basePath, trs)
	if storedTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedTrsErr.Error())
	}

	//retrieve the transaction:
	retTrs, retTrsErr := repository.Retrieve(basePath)
	if retTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retTrsErr.Error())
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the retrieved transaction is invalid")
	}

	//delete the transaction:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
