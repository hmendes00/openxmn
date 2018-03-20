package transactions

import (
	"path/filepath"
	"reflect"
	"testing"

	transactions "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions"
	conncrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
)

func TestSave_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateTransactionForTests()
	secondTrs := CreateTransactionForTests()
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

	//factories:
	fileService := conncrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	trsRepository := CreateTransactionRepositoryForTests()
	trsService := CreateTransactionServiceForTests()

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
