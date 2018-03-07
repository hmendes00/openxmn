package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	concrete_stored_users "github.com/XMNBlockchain/core/packages/storages/users/infrastructure"
)

func TestSaveUser_thenRetrieve_Success(t *testing.T) {

	//create the user:
	usr := CreateUserForTests(t)

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	pubKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	fileRepository := concrete_files.CreateFileRepository(fileBuilderFactory)
	fileService := concrete_files.CreateFileService(storedFileBuilderFactory)
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	metaDataRepository := concrete_metadata.CreateMetaDataRepository(fileRepository)
	metaDataService := concrete_metadata.CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	usrBuilderFactory := CreateUserBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedUserBuilderFactory := concrete_stored_users.CreateUserBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateUserRepository(metaDataRepository, fileRepository, pubKeyBuilderFactory, usrBuilderFactory)
	service := CreateUserService(metaDataService, fileService, fileBuilderFactory, storedUserBuilderFactory)

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
