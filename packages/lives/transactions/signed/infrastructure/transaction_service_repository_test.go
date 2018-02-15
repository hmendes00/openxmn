package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	conncrete_objects "github.com/XMNBlockchain/core/packages/lives/objects/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	conncrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/infrastructure"
)

func TestSaveTrs_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateTransactionForTests(t)
	secondTrs := CreateTransactionForTests(t)
	multipleTrs := []signed_transactions.Transaction{
		trs,
		secondTrs,
	}

	multipleTrsMap := map[string]signed_transactions.Transaction{
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

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateTransactionRepository(objectRepository, trsRepository, signedTrsBuilderFactory)
	service := CreateTransactionService(metaDataBuilderFactory, storedTreeBuilderFactory, trsService, objectService, objBuilderFactory)

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
		t.Errorf("the amount of retrieved signed transactions is invalid.  Expected: %d, Received: %d", len(multipleTrs), len(retMultipleTrs))
	}

	for index, oneRetTrs := range retMultipleTrs {
		retIDAsString := oneRetTrs.GetID().String()
		if oneTrs, ok := multipleTrsMap[retIDAsString]; ok {
			if !reflect.DeepEqual(oneTrs, oneRetTrs) {
				t.Errorf("the retrieved signed transaction at index: %d (ID: %s) is invalid", index, retIDAsString)
			}

			continue
		}

		t.Errorf("the retrieved signed transaction (ID: %s) should not exists", retIDAsString)
	}
}
