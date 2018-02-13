package infrastructure

import (
	"fmt"
	"path/filepath"
	"reflect"
	"testing"

	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	conncrete_objects "github.com/XMNBlockchain/core/packages/lives/objects/infrastructure"
	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/infrastructure"
)

func TestSave_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateTransactionForTests(t)
	secondTrs := CreateTransactionForTests(t)
	multipleTrs := []transactions.Transaction{
		trs,
		secondTrs,
	}

	multipleTrsMap := map[string]transactions.Transaction{
		trs.GetID().String():       trs,
		secondTrs.GetID().String(): secondTrs,
	}

	//variables:
	basePath := filepath.Join("test_files", "transactions")
	lastDirPath := fmt.Sprintf("%s_%d", trs.GetID().String(), trs.CreatedOn().UnixNano())
	basePathWithTrs := filepath.Join(basePath, lastDirPath)
	chkSizeInBytes := 8
	extension := "chk"

	//factories:
	objBuilderFactory := conncrete_objects.CreateObjectBuilderFactory()
	fileBuilderFactory := conncrete_files.CreateFileBuilderFactory()
	fileRepository := conncrete_files.CreateFileRepository(fileBuilderFactory)
	htBuilderFactory := conncrete_hashtrees.CreateHashTreeBuilderFactory()
	chksBuilderFactory := conncrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := conncrete_chunks.CreateChunksRepository(fileRepository, chksBuilderFactory)
	storedFileBuilderFactory := conncrete_stored_files.CreateFileBuilderFactory()
	fileService := conncrete_files.CreateFileService(storedFileBuilderFactory)
	storedChkBuilderFactory := conncrete_stored_chunks.CreateChunksBuilderFactory()
	chkService := conncrete_chunks.CreateChunksService(fileService, fileBuilderFactory, storedChkBuilderFactory)
	storedObjBuilderFactory := conncrete_stored_objects.CreateObjectBuilderFactory()
	objRepository := conncrete_objects.CreateObjectRepository(objBuilderFactory, chkRepository, fileRepository)
	objService := conncrete_objects.CreateObjectService(storedObjBuilderFactory, fileBuilderFactory, fileService, chkService, htBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	trsRepository := CreateTransactionRepository(objRepository)
	trsService := CreateTransactionService(objService, chksBuilderFactory, objBuilderFactory, storedObjBuilderFactory)

	//make sure there is no transactions:
	_, noTrsErr := trsRepository.RetrieveAll(basePath)
	if noTrsErr == nil {
		t.Errorf("there was supposed to be no transaction.")
	}

	//save the transaction:
	_, storedTrsErr := trsService.Save(basePath, trs)
	if storedTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedTrsErr.Error())
	}

	//retrieve the transaction:
	retTrs, retTrsErr := trsRepository.Retrieve(basePathWithTrs)
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
	_, multipleSaveErr := trsService.SaveAll(basePath, multipleTrs)
	if multipleSaveErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", multipleSaveErr.Error())
	}

	//retrieve multiple trs:
	retMultipleTrs, retMultipleTrsErr := trsRepository.RetrieveAll(basePath)
	if retMultipleTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retMultipleTrsErr.Error())
	}

	if len(retMultipleTrs) != len(multipleTrs) {
		t.Errorf("the amount of retrieved transactions is invalid.  Expected: %d, Received: %d", len(multipleTrs), len(retMultipleTrs))
	}

	for index, oneRetTrs := range retMultipleTrs {
		retIDAsString := oneRetTrs.GetID().String()
		if oneTrs, ok := multipleTrsMap[retIDAsString]; ok {
			if !reflect.DeepEqual(oneTrs, oneRetTrs) {
				t.Errorf("the retrieved transaction at index: %d (ID: %s) is invalid", index, retIDAsString)
			}

			continue
		}

		t.Errorf("the retrieved transaction (ID: %s) should not exists", retIDAsString)
	}
}
