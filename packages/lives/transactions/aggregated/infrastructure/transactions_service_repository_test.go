package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	conncrete_objects "github.com/XMNBlockchain/core/packages/lives/objects/infrastructure"
	concrete_signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
	conncrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/infrastructure"
)

func TestSaveAtomicTrs_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateTransactionsForTests(t)

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 8
	extension := "chk"

	//factories:
	objBuilderFactory := conncrete_objects.CreateObjectBuilderFactory()
	fileBuilderFactory := conncrete_files.CreateFileBuilderFactory()
	storedFileBuilderFactory := conncrete_stored_files.CreateFileBuilderFactory()
	storedChkBuilderFactory := conncrete_stored_chunks.CreateChunksBuilderFactory()
	fileRepository := conncrete_files.CreateFileRepository(fileBuilderFactory)
	fileService := conncrete_files.CreateFileService(storedFileBuilderFactory)
	htBuilderFactory := conncrete_hashtrees.CreateHashTreeBuilderFactory()
	htService := conncrete_hashtrees.CreateHashTreeService(fileService, fileBuilderFactory)
	htRepository := conncrete_hashtrees.CreateHashTreeRepository(fileRepository)
	objsBuilderFactory := conncrete_objects.CreateObjectsBuilderFactory(htBuilderFactory)
	chksBuilderFactory := conncrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := conncrete_chunks.CreateChunksRepository(htRepository, fileRepository, chksBuilderFactory)
	chkService := conncrete_chunks.CreateChunksService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	storedObjBuilderFactory := conncrete_stored_objects.CreateObjectBuilderFactory()
	metaDataBuilderFactory := conncrete_objects.CreateMetaDataBuilderFactory()
	metaDataRepository := conncrete_objects.CreateMetaDataRepository(fileRepository)
	metaDataService := conncrete_objects.CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	objectRepository := conncrete_objects.CreateObjectRepository(metaDataRepository, objBuilderFactory, chkRepository)
	objectService := conncrete_objects.CreateObjectService(metaDataService, storedObjBuilderFactory, chkService)
	trsRepository := conncrete_transactions.CreateTransactionRepository(objectRepository)
	trsService := conncrete_transactions.CreateTransactionService(objectService, metaDataBuilderFactory, chksBuilderFactory, objBuilderFactory, storedObjBuilderFactory)
	signedTrsBuilderFactory := concrete_signed_transactions.CreateTransactionBuilderFactory()
	storedTreeBuilderFactory := conncrete_stored_objects.CreateTreeBuilderFactory()
	storedTreesBuilderFactory := conncrete_stored_objects.CreateTreesBuilderFactory()
	atomicTrsBuilderFactory := concrete_signed_transactions.CreateAtomicTransactionBuilderFactory(htBuilderFactory)
	storedObjsBuilderFactory := conncrete_stored_objects.CreateObjectsBuilderFactory()
	signedTrsRepository := concrete_signed_transactions.CreateTransactionRepository(objectRepository, trsRepository, signedTrsBuilderFactory)
	signedTrsService := concrete_signed_transactions.CreateTransactionService(metaDataBuilderFactory, storedTreeBuilderFactory, trsService, objectService, objBuilderFactory)
	atomicTrsRepository := concrete_signed_transactions.CreateAtomicTransactionRepository(objectRepository, fileRepository, trsRepository, signedTrsBuilderFactory, atomicTrsBuilderFactory)
	atomicTrsService := concrete_signed_transactions.CreateAtomicTransactionService(metaDataBuilderFactory, fileBuilderFactory, fileService, storedTreeBuilderFactory, trsService, objectService, objBuilderFactory, objsBuilderFactory, storedObjsBuilderFactory)
	aggregatedTrsBuilderFactory := CreateTransactionsBuilderFactory(htBuilderFactory)
	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateTransactionsRepository(signedTrsRepository, atomicTrsRepository, htRepository, objectRepository, aggregatedTrsBuilderFactory)
	service := CreateTransactionsService(signedTrsService, atomicTrsService, htService, metaDataBuilderFactory, objectService, objBuilderFactory, storedTreeBuilderFactory, storedTreesBuilderFactory)

	//make sure there is no transactions:
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
