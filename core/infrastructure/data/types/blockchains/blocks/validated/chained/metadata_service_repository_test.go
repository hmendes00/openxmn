package chained

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/files"
)

func TestSaveMetaData_thenRetrieve_Success(t *testing.T) {

	//create the block:
	met := CreateMetaDataForTests()

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateMetaDataRepositoryForTests()
	service := CreateMetaDataServiceForTests()

	//make sure there is no metadata:
	_, noMetErr := repository.Retrieve(basePath)
	if noMetErr == nil {
		t.Errorf("there was supposed to be no metadata.")
	}

	//save the metadata:
	_, storedMetErr := service.Save(basePath, met)
	if storedMetErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedMetErr.Error())
	}

	//retrieve the metadata:
	retMet, retMetErr := repository.Retrieve(basePath)
	if retMetErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retMetErr.Error())
	}

	if !reflect.DeepEqual(met, retMet) {
		t.Errorf("the retrieved metadata is invalid")
	}

	//delete the metadata:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
