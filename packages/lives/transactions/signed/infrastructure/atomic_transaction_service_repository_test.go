package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	conncrete_objects "github.com/XMNBlockchain/core/packages/lives/objects/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	conncrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/infrastructure"
)

func TestSaveAtomicTrs_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateAtomicTransactionForTests(t)
	secondTrs := CreateAtomicTransactionForTests(t)
	multipleTrs := []signed_transactions.AtomicTransaction{
		trs,
		secondTrs,
	}

	multipleTrsMap := map[string]signed_transactions.AtomicTransaction{
		trs.GetID().String():       trs,
		secondTrs.GetID().String(): secondTrs,
	}

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
	signedTrsBuilderFactory := CreateTransactionBuilderFactory()
	storedTreeBuilderFactory := conncrete_stored_objects.CreateTreeBuilderFactory()
	atomicTrsBuilderFactory := CreateAtomicTransactionBuilderFactory(htBuilderFactory)
	storedObjsBuilderFactory := conncrete_stored_objects.CreateObjectsBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateAtomicTransactionRepository(objectRepository, fileRepository, trsRepository, signedTrsBuilderFactory, atomicTrsBuilderFactory)
	service := CreateAtomicTransactionService(metaDataBuilderFactory, fileBuilderFactory, fileService, storedTreeBuilderFactory, trsService, objectService, objBuilderFactory, objsBuilderFactory, storedObjsBuilderFactory)

	//make sure there is no transactions:
	_, noTrsErr := repository.RetrieveAll(basePath)
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

	//save multiple transactions:
	_, multipleSaveErr := service.SaveAll(basePath, multipleTrs)
	if multipleSaveErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", multipleSaveErr.Error())
	}

	//retrieve multiple trs:
	retMultipleTrs, retMultipleTrsErr := repository.RetrieveAll(basePath)
	if retMultipleTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retMultipleTrsErr.Error())
	}

	if len(retMultipleTrs) != len(multipleTrs) {
		t.Errorf("the amount of retrieved atomic transactions is invalid.  Expected: %d, Received: %d", len(multipleTrs), len(retMultipleTrs))
	}

	for index, oneRetTrs := range retMultipleTrs {
		retIDAsString := oneRetTrs.GetID().String()
		if oneTrs, ok := multipleTrsMap[retIDAsString]; ok {
			if !reflect.DeepEqual(oneTrs, oneRetTrs) {
				t.Errorf("the retrieved atomic transaction at index: %d (ID: %s) is invalid", index, retIDAsString)
			}

			continue
		}

		t.Errorf("the retrieved atomic transaction (ID: %s) should not exists", retIDAsString)
	}
}
