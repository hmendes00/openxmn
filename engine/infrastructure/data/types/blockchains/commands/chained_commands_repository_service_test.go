package commands

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
)

func TestSaveChainedCommands_thenRetrieve_Success(t *testing.T) {

	//create the chained commands:
	chainedCommands := CreateChainedCommandsForTests()

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateChainedCommandsRepositoryForTests()
	service := CreateChainedCommandsServiceForTests()

	//make sure there is no chained commands:
	_, noChainedCmdsErr := repository.Retrieve(basePath)
	if noChainedCmdsErr == nil {
		t.Errorf("there was supposed to be no chained commands.")
	}

	//save the chained commands:
	_, storedChainedCmdsErr := service.Save(basePath, chainedCommands)
	if storedChainedCmdsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedChainedCmdsErr.Error())
	}

	//retrieve the chained commands:
	retChainedCommands, retChainedCommandsErr := repository.Retrieve(basePath)
	if retChainedCommandsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retChainedCommandsErr.Error())
	}

	if !reflect.DeepEqual(chainedCommands, retChainedCommands) {
		t.Errorf("the retrieved chained commands is invalid")
	}

	//delete the chained commands:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
