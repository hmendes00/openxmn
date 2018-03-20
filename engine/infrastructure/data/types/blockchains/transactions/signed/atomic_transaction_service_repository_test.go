package signed

import (
	"path/filepath"
	"reflect"
	"testing"

	signed_transactions "github.com/XMNBlockchain/exmachina-network/engine/domain/data/types/blockchains/transactions/signed"
	concrete_files "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/types/files"
)

func TestSaveAtomicTrs_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateAtomicTransactionForTests()
	secondTrs := CreateAtomicTransactionForTests()
	multipleTrs := []signed_transactions.AtomicTransaction{
		trs,
		secondTrs,
	}

	multipleTrsMap := map[string]signed_transactions.AtomicTransaction{
		trs.GetMetaData().GetID().String():       trs,
		secondTrs.GetMetaData().GetID().String(): secondTrs,
	}

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateAtomicTransactionRepositoryForTests()
	service := CreateAtomicTransactionServiceForTests()

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
		retIDAsString := oneRetTrs.GetMetaData().GetID().String()
		if oneTrs, ok := multipleTrsMap[retIDAsString]; ok {
			if !reflect.DeepEqual(oneTrs, oneRetTrs) {
				t.Errorf("the retrieved atomic transaction at index: %d (ID: %s) is invalid", index, retIDAsString)
			}

			continue
		}

		t.Errorf("the retrieved atomic transaction (ID: %s) should not exists", retIDAsString)
	}
}
