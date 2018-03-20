package users

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/files"
)

func TestSaveUser_thenRetrieve_Success(t *testing.T) {

	//create the user:
	usr := CreateUserForTests()

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateUserRepositoryForTests()
	service := CreateUserServiceForTests()

	//make sure there is no user:
	_, noUserErr := repository.Retrieve(basePath)
	if noUserErr == nil {
		t.Errorf("there was supposed to be no user.")
	}

	//save the user:
	_, storedUserErr := service.Save(basePath, usr)
	if storedUserErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedUserErr.Error())
	}

	//retrieve the user:
	retUser, retUserErr := repository.Retrieve(basePath)
	if retUserErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retUserErr.Error())
	}

	if !reflect.DeepEqual(usr, retUser) {
		t.Errorf("the retrieved user is invalid")
	}

	//delete the user:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
