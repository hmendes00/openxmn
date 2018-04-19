package commands

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
)

func TestSaveCommands_thenRetrieve_Success(t *testing.T) {

	//create the commands:
	commands := CreateCommandsForTests()

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateRepositoryForTests()
	service := CreateServiceForTests()

	//make sure there is no commands:
	_, noChainedCmdsErr := repository.Retrieve(basePath)
	if noChainedCmdsErr == nil {
		t.Errorf("there was supposed to be no commands.")
	}

	//save the commands:
	_, storedChainedCmdsErr := service.Save(basePath, commands)
	if storedChainedCmdsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedChainedCmdsErr.Error())
	}

	//retrieve the commands:
	retCommands, retCommandsErr := repository.Retrieve(basePath)
	if retCommandsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retCommandsErr.Error())
	}

	if !reflect.DeepEqual(commands, retCommands) {
		t.Errorf("the retrieved commands is invalid")
	}

	//delete the commands:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
