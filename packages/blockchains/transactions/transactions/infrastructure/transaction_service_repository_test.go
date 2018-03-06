package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	conncrete_chunks "github.com/XMNBlockchain/core/packages/blockchains/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	conncrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/domain"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/infrastructure"
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
		trs.GetMetaData().GetID().String():       trs,
		secondTrs.GetMetaData().GetID().String(): secondTrs,
	}

	//variables:
	basePath := filepath.Join("test_files", "transactions")
	chkSizeInBytes := 8
	extension := "chk"

	//factories:
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
	metaDataService := conncrete_metadata.CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	metaDataRepository := conncrete_metadata.CreateMetaDataRepository(fileRepository)
	metBuilderFactory := conncrete_metadata.CreateMetaDataBuilderFactory()
	storedTrsBuilderFactory := conncrete_stored_transactions.CreateTransactionBuilderFactory()
	trsBuilderFactory := CreateTransactionBuilderFactory(htBuilderFactory, metBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	trsRepository := CreateTransactionRepository(chkRepository, metaDataRepository, trsBuilderFactory)
	trsService := CreateTransactionService(metaDataService, chksBuilderFactory, chkService, storedTrsBuilderFactory)

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
	retTrs, retTrsErr := trsRepository.Retrieve(basePath)
	if retTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retTrsErr.Error())
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the retrieved transaction JSON is invalid.")
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
		retIDAsString := oneRetTrs.GetMetaData().GetID().String()
		if oneTrs, ok := multipleTrsMap[retIDAsString]; ok {
			if !reflect.DeepEqual(oneTrs, oneRetTrs) {
				t.Errorf("the retrieved transaction at index: %d (ID: %s) is invalid", index, retIDAsString)
			}

			continue
		}

		t.Errorf("the retrieved transaction (ID: %s) should not exists", retIDAsString)
	}
}
