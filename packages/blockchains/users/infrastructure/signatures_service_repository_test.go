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

func TestSaveUserSignatures_thenRetrieve_Success(t *testing.T) {

	//create the signature:
	sigs := CreateSignaturesForTests(t)

	//variables:
	basePath := filepath.Join("test_files", "files")

	//factories:
	pubKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(pubKeyBuilderFactory)
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
	usrRepository := CreateUserRepository(metaDataRepository, fileRepository, pubKeyBuilderFactory, usrBuilderFactory)
	usrService := CreateUserService(metaDataService, fileService, fileBuilderFactory, storedUserBuilderFactory)
	userSigBuilderFactory := CreateSignatureBuilderFactory(sigBuilderFactory, htBuilderFactory, metaDataBuilderFactory)
	storedSigBuilderFactory := concrete_stored_users.CreateSignatureBuilderFactory()
	userSigRepository := CreateSignatureRepository(metaDataRepository, usrRepository, fileRepository, userSigBuilderFactory)
	userSigService := CreateSignatureService(metaDataService, usrService, fileService, fileBuilderFactory, storedSigBuilderFactory)
	storedSigsBuilderFactory := concrete_stored_users.CreateSignaturesBuilderFactory()
	sigsBuilderFactory := CreateSignaturesBuilderFactory(htBuilderFactory, metaDataBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateSignaturesRepository(metaDataRepository, userSigRepository, sigsBuilderFactory)
	service := CreateSignaturesService(metaDataService, userSigService, storedSigsBuilderFactory)

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

	if !reflect.DeepEqual(sigs, retSigs) {
		t.Errorf("the retrieved signatures is invalid")
	}

	//delete the sig:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
