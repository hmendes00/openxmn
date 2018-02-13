package infrastructure

import (
	"path/filepath"
	"testing"

	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	conncrete_objects "github.com/XMNBlockchain/core/packages/lives/objects/infrastructure"
	conncrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/infrastructure"
)

func TestSave_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateTransactionForTests(t)
	//	secondTrs := CreateTransactionForTests(t)
	/*multipleTrs := []signed_transactions.Transaction{
		trs,
		secondTrs,
	}
	*/
	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 8
	extension := "chk"

	//factories:
	objBuilderFactory := conncrete_objects.CreateObjectBuilderFactory()
	fileBuilderFactory := conncrete_files.CreateFileBuilderFactory()
	//fileRepository := conncrete_files.CreateFileRepository(fileBuilderFactory)
	htBuilderFactory := conncrete_hashtrees.CreateHashTreeBuilderFactory()
	chksBuilderFactory := conncrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	//chkRepository := conncrete_chunks.CreateChunksRepository(fileRepository, chksBuilderFactory)
	storedFileBuilderFactory := conncrete_stored_files.CreateFileBuilderFactory()
	fileService := conncrete_files.CreateFileService(storedFileBuilderFactory)
	storedChkBuilderFactory := conncrete_stored_chunks.CreateChunksBuilderFactory()
	chkService := conncrete_chunks.CreateChunksService(fileService, fileBuilderFactory, storedChkBuilderFactory)
	storedObjBuilderFactory := conncrete_stored_objects.CreateObjectBuilderFactory()
	//objRepository := conncrete_objects.CreateObjectRepository(objBuilderFactory, chkRepository, fileRepository)
	objService := conncrete_objects.CreateObjectService(storedObjBuilderFactory, fileBuilderFactory, fileService, chkService, htBuilderFactory)
	//trsRepository := conncrete_transactions.CreateTransactionRepository(objRepository)
	trsService := conncrete_transactions.CreateTransactionService(objService, chksBuilderFactory, objBuilderFactory, storedObjBuilderFactory)
	//signedTrsBuilderFactory := CreateTransactionBuilderFactory()
	storedTreeBuilderFactory := conncrete_stored_objects.CreateTreeBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		//fileService.DeleteAll(basePath)
	}()

	//execute:
	//repository := CreateTransactionRepository(objRepository, trsRepository, signedTrsBuilderFactory)
	service := CreateTransactionService(storedTreeBuilderFactory, trsService, objService, objBuilderFactory)

	//make sure there is no transactions:
	/*	_, noTrsErr := repository.RetrieveAll(basePath)
		if noTrsErr == nil {
			t.Errorf("there was supposed to be no transaction.")
		}*/

	//save the transaction:
	_, storedTrsErr := service.Save(basePath, trs)
	if storedTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedTrsErr.Error())
	}

	//retrieve the transaction:
	/*retTrs, retTrsErr := repository.Retrieve(basePath)
	if retTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retTrsErr.Error())
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the retrieved transaction is invalid")
	}*/

	//delete the transaction:
	/*	delErr := fileService.DeleteAll(basePath)
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

		if !reflect.DeepEqual(multipleTrs, retMultipleTrs) {
			t.Errorf("the retrieved []Transaction are invalid")
		}*/
}
