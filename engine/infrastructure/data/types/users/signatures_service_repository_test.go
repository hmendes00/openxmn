package users

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
)

func TestSaveUserSignatures_thenRetrieve_Success(t *testing.T) {

	//create the signature:
	sigs := CreateSignaturesForTests()

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateSignaturesRepositoryForTests()
	service := CreateSignaturesServiceForTests()

	//make sure there is no sigs:
	_, noSigErr := repository.Retrieve(basePath)
	if noSigErr == nil {
		t.Errorf("there was supposed to be no signature.")
	}

	//save the sigs:
	_, storedSigErr := service.Save(basePath, sigs)
	if storedSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedSigErr.Error())
	}

	//retrieve the sigs:
	retSigs, retSigErr := repository.Retrieve(basePath)
	if retSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retSigErr.Error())
	}

	if !reflect.DeepEqual(retSigs.GetMetaData(), sigs.GetMetaData()) {
		t.Errorf("the signatures metadata was invalid")
	}

	//the returned signatures can be in another order:
	for _, oneRetSig := range retSigs.GetSignatures() {
		isFound := false
		for _, oneSig := range sigs.GetSignatures() {
			if reflect.DeepEqual(oneRetSig, oneSig) {
				isFound = true
				continue
			}
		}

		if !isFound {
			t.Errorf("the returned signature (ID: %s) could not be found in the original signatures", oneRetSig.GetMetaData().GetID())
		}

	}

	//delete the sig:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
