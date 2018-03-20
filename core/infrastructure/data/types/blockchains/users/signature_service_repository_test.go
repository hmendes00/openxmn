package users

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/files"
)

func TestSaveUserSignature_thenRetrieve_Success(t *testing.T) {

	//create the signature:
	sig := CreateSignatureForTests()

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateSignatureRepositoryForTests()
	service := CreateSignatureServiceForTests()

	//make sure there is no sig:
	_, noSigErr := repository.Retrieve(basePath)
	if noSigErr == nil {
		t.Errorf("there was supposed to be no signature.")
	}

	//save the sig:
	_, storedSigErr := service.Save(basePath, sig)
	if storedSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedSigErr.Error())
	}

	//retrieve the sig:
	retSig, retSigErr := repository.Retrieve(basePath)
	if retSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retSigErr.Error())
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the retrieved signature is invalid")
	}

	//delete the sig:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
