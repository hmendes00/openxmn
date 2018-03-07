package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_chunks "github.com/XMNBlockchain/core/packages/blockchains/chunks/infrastructure"
	concrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	concrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/infrastructure"
	concrete_stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/infrastructure"
	concrete_stored_users "github.com/XMNBlockchain/core/packages/storages/users/infrastructure"
)

func TestSaveTransaction_thenRetrieve_Success(t *testing.T) {

	//create the transaction:
	trs := CreateTransactionForTests(t)
	secondTrs := CreateTransactionForTests(t)
	multipleTrs := []signed_transactions.Transaction{
		trs,
		secondTrs,
	}

	multipleTrsMap := map[string]signed_transactions.Transaction{
		trs.GetMetaData().GetID().String():       trs,
		secondTrs.GetMetaData().GetID().String(): secondTrs,
	}

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 16
	extension := "chk"

	//factories:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	storedChkBuilderFactory := concrete_stored_chunks.CreateChunksBuilderFactory()
	fileRepository := concrete_files.CreateFileRepository(fileBuilderFactory)
	fileService := concrete_files.CreateFileService(storedFileBuilderFactory)
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	htService := concrete_hashtrees.CreateHashTreeService(fileService, fileBuilderFactory)
	htRepository := concrete_hashtrees.CreateHashTreeRepository(fileRepository)
	chksBuilderFactory := concrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := concrete_chunks.CreateChunksRepository(htRepository, fileRepository, chksBuilderFactory)
	chkService := concrete_chunks.CreateChunksService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	metaDataRepository := concrete_metadata.CreateMetaDataRepository(fileRepository)
	metaDataService := concrete_metadata.CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	trsBuilderFactory := concrete_transactions.CreateTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	trsRepository := concrete_transactions.CreateTransactionRepository(chkRepository, metaDataRepository, trsBuilderFactory)
	storedTrsBuilderFactory := concrete_stored_transactions.CreateTransactionBuilderFactory()
	trsService := concrete_transactions.CreateTransactionService(metaDataService, chksBuilderFactory, chkService, storedTrsBuilderFactory)
	signedTrsBuilderFactory := CreateTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	pubKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(pubKeyBuilderFactory)
	usrBuilderFactory := concrete_users.CreateUserBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	usrRepository := concrete_users.CreateUserRepository(metaDataRepository, fileRepository, pubKeyBuilderFactory, usrBuilderFactory)
	storedUserBuilderFactory := concrete_stored_users.CreateUserBuilderFactory()
	usrService := concrete_users.CreateUserService(metaDataService, fileService, fileBuilderFactory, storedUserBuilderFactory)
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, htBuilderFactory, metaDataBuilderFactory)
	storedSigBuilderFactory := concrete_stored_users.CreateSignatureBuilderFactory()
	userSigRepository := concrete_users.CreateSignatureRepository(metaDataRepository, usrRepository, fileRepository, userSigBuilderFactory)
	userSigService := concrete_users.CreateSignatureService(metaDataService, usrService, fileService, fileBuilderFactory, storedSigBuilderFactory)
	storedSignedTrsBuilderFactory := concrete_stored_signed_transactions.CreateTransactionBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateTransactionRepository(metaDataRepository, userSigRepository, trsRepository, signedTrsBuilderFactory)
	service := CreateTransactionService(metaDataService, trsService, storedSignedTrsBuilderFactory, userSigService)

	//make sure there is no transactions:
	_, noTrsErr := repository.RetrieveAll(basePath)
	if noTrsErr == nil {
		t.Errorf("there was supposed to be no transaction.")
	}

	//save the transaction:
	_, storedTrsErr := service.Save(basePath, trs)
	if storedTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedTrsErr.Error())
	}

	//retrieve the transaction:
	retTrs, retTrsErr := repository.Retrieve(basePath)
	if retTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retTrsErr.Error())
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the retrieved transaction is invalid")
	}

	//delete the transaction:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}

	//save multiple transactions:
	_, multipleSaveErr := service.SaveAll(basePath, multipleTrs)
	if multipleSaveErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", multipleSaveErr.Error())
	}

	//retrieve multiple trs:
	retMultipleTrs, retMultipleTrsErr := repository.RetrieveAll(basePath)
	if retMultipleTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retMultipleTrsErr.Error())
	}

	if len(retMultipleTrs) != len(multipleTrs) {
		t.Errorf("the amount of retrieved signed transactions is invalid.  Expected: %d, Received: %d", len(multipleTrs), len(retMultipleTrs))
	}

	for index, oneRetTrs := range retMultipleTrs {
		retIDAsString := oneRetTrs.GetMetaData().GetID().String()
		if oneTrs, ok := multipleTrsMap[retIDAsString]; ok {
			if !reflect.DeepEqual(oneTrs, oneRetTrs) {
				t.Errorf("the retrieved signed transaction at index: %d (ID: %s) is invalid", index, retIDAsString)
			}

			continue
		}

		t.Errorf("the retrieved signed transaction (ID: %s) should not exists", retIDAsString)
	}
}
