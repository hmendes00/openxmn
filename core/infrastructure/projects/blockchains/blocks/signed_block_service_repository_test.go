package blocks

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/files"
)

func TestSaveSignedBlock_thenRetrieve_Success(t *testing.T) {

	//create the block:
	signedBlk := CreateSignedBlockForTests()

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateSignedBlockRepositoryForTests()
	service := CreateSignedBlockServiceForTests()

	//make sure there is no blocks:
	_, noTrsErr := repository.Retrieve(basePath)
	if noTrsErr == nil {
		t.Errorf("there was supposed to be no signed block.")
	}

	//save the block:
	_, storedTrsErr := service.Save(basePath, signedBlk)
	if storedTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedTrsErr.Error())
	}

	//retrieve the block:
	retBlk, retBlkErr := repository.Retrieve(basePath)
	if retBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retBlkErr.Error())
	}

	if !reflect.DeepEqual(signedBlk, retBlk) {
		t.Errorf("the retrieved block is invalid")
	}

	//delete the block:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
