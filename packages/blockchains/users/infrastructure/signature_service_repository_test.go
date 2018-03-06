package infrastructure

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/domain"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
)

func TestSaveUserSignature_thenRetrieve_Success(t *testing.T) {

	//signature:
	sig := CreateSignatureForTests(t)
	secondSig := CreateSignatureForTests(t)
	sigs := []users.Signature{
		sig,
		secondSig,
	}

	sigsMap := map[string]users.Signature{
		sig.GetMetaData().GetID().String():       sig,
		secondSig.GetMetaData().GetID().String(): secondSig,
	}

	//file variables:
	saveInPath := filepath.Join("test_files", "files")

	//factories:
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	fileService := concrete_files.CreateFileService(storedFileBuilderFactory)
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	fileRepository := concrete_files.CreateFileRepository(fileBuilderFactory)

	//delete the files at the end:
	defer func() {
		fileService.DeleteAll(saveInPath)
	}()

	//execute:
	service := CreateSignatureService(fileService, fileBuilderFactory)
	repository := CreateSignatureRepository(fileRepository)

	//verify that the file do not exists:
	if _, err := os.Stat(saveInPath); !os.IsNotExist(err) {
		t.Errorf("the given path was not expected to be a valid directory: %s", saveInPath)
	}

	//save the signature on disk:
	_, storedSigErr := service.Save(saveInPath, sig)
	if storedSigErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedSigErr.Error())
	}

	//retrieve the signature:
	retSig, retSigErr := repository.Retrieve(saveInPath)
	if retSigErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retSigErr.Error())
	}

	//compare the signatures:
	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned signature is invalid")
	}

	//delete the files:
	delErr := fileService.DeleteAll(saveInPath)
	if delErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", delErr.Error())
	}

	//save the signatures:
	_, storedFilesErr := service.SaveAll(saveInPath, sigs)
	if storedFilesErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedFilesErr.Error())
	}

	//retrieve the signatures:
	retSigs, retSigsErr := repository.RetrieveAll(saveInPath)
	if retSigsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retSigsErr.Error())
	}

	//compare the signatures:
	for index, oneRetSig := range retSigs {
		idAsString := oneRetSig.GetMetaData().GetID().String()
		if foundSig, ok := sigsMap[idAsString]; ok {
			if !reflect.DeepEqual(foundSig, oneRetSig) {
				t.Errorf("the returned signature (index: %d, ID: %s) is invalid", index, idAsString)
			}

			continue
		}

		t.Errorf("the signature (ID: %s) could not be found", idAsString)
	}
}
