package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
)

func TestSaveMetaData_thenRetrieve_Success(t *testing.T) {

	//create the block:
	met := CreateMetaDataForTests(t)

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	fileRepository := concrete_files.CreateFileRepository(fileBuilderFactory)
	fileService := concrete_files.CreateFileService(storedFileBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateMetaDataRepository(fileRepository)
	service := CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)

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
