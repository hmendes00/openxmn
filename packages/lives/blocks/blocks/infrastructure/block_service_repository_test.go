package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	conncrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	conncrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	conncrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	conncrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	conncrete_objects "github.com/XMNBlockchain/core/packages/lives/objects/infrastructure"
	conncrete_aggregated_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
	concrete_signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
	conncrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	conncrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	conncrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	conncrete_stored_objects "github.com/XMNBlockchain/core/packages/storages/objects/infrastructure"
	conncrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
)

func TestSaveBlock_thenRetrieve_Success(t *testing.T) {

	//create the block:
	blk := CreateBlockForTests(t)

	//variables:
	basePath := filepath.Join("test_files", "files")
	chkSizeInBytes := 8
	extension := "chk"

	//factories:
	objBuilderFactory := conncrete_objects.CreateObjectBuilderFactory()
	fileBuilderFactory := conncrete_files.CreateFileBuilderFactory()
	storedFileBuilderFactory := conncrete_stored_files.CreateFileBuilderFactory()
	storedChkBuilderFactory := conncrete_stored_chunks.CreateChunksBuilderFactory()
	fileRepository := conncrete_files.CreateFileRepository(fileBuilderFactory)
	fileService := conncrete_files.CreateFileService(storedFileBuilderFactory)
	htBuilderFactory := conncrete_hashtrees.CreateHashTreeBuilderFactory()
	htService := conncrete_hashtrees.CreateHashTreeService(fileService, fileBuilderFactory)
	htRepository := conncrete_hashtrees.CreateHashTreeRepository(fileRepository)
	objsBuilderFactory := conncrete_objects.CreateObjectsBuilderFactory(htBuilderFactory)
	chksBuilderFactory := conncrete_chunks.CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	chkRepository := conncrete_chunks.CreateChunksRepository(htRepository, fileRepository, chksBuilderFactory)
	chkService := conncrete_chunks.CreateChunksService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	storedObjBuilderFactory := conncrete_stored_objects.CreateObjectBuilderFactory()
	metaDataBuilderFactory := conncrete_objects.CreateMetaDataBuilderFactory()
	metaDataRepository := conncrete_objects.CreateMetaDataRepository(fileRepository)
	metaDataService := conncrete_objects.CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	objectRepository := conncrete_objects.CreateObjectRepository(metaDataRepository, objBuilderFactory, chkRepository)
	objectService := conncrete_objects.CreateObjectService(metaDataService, storedObjBuilderFactory, chkService)
	trsRepository := conncrete_transactions.CreateTransactionRepository(objectRepository)
	trsService := conncrete_transactions.CreateTransactionService(objectService, metaDataBuilderFactory, chksBuilderFactory, objBuilderFactory, storedObjBuilderFactory)
	signedTrsBuilderFactory := concrete_signed_transactions.CreateTransactionBuilderFactory()
	storedTreeBuilderFactory := conncrete_stored_objects.CreateTreeBuilderFactory()
	storedTreesBuilderFactory := conncrete_stored_objects.CreateTreesBuilderFactory()
	atomicTrsBuilderFactory := concrete_signed_transactions.CreateAtomicTransactionBuilderFactory(htBuilderFactory)
	storedObjsBuilderFactory := conncrete_stored_objects.CreateObjectsBuilderFactory()
	signedTrsRepository := concrete_signed_transactions.CreateTransactionRepository(objectRepository, trsRepository, signedTrsBuilderFactory)
	signedTrsService := concrete_signed_transactions.CreateTransactionService(metaDataBuilderFactory, storedTreeBuilderFactory, trsService, objectService, objBuilderFactory)
	atomicTrsRepository := concrete_signed_transactions.CreateAtomicTransactionRepository(objectRepository, fileRepository, trsRepository, signedTrsBuilderFactory, atomicTrsBuilderFactory)
	atomicTrsService := concrete_signed_transactions.CreateAtomicTransactionService(metaDataBuilderFactory, fileBuilderFactory, fileService, storedTreeBuilderFactory, trsService, objectService, objBuilderFactory, objsBuilderFactory, storedObjsBuilderFactory)
	aggregatedTrsBuilderFactory := conncrete_aggregated_transactions.CreateTransactionsBuilderFactory(htBuilderFactory)
	aggregatedTrsRepository := conncrete_aggregated_transactions.CreateTransactionsRepository(signedTrsRepository, atomicTrsRepository, htRepository, objectRepository, aggregatedTrsBuilderFactory)
	aggregatedTrsService := conncrete_aggregated_transactions.CreateTransactionsService(signedTrsService, atomicTrsService, htService, metaDataBuilderFactory, objectService, objBuilderFactory, storedTreeBuilderFactory, storedTreesBuilderFactory)
	publicKeyBuilderFactory := conncrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := conncrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userBuilderFactory := conncrete_users.CreateUserBuilderFactory()
	userSigBuilderFactory := conncrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, userBuilderFactory)
	aggregatedSignedTrsBuilderFactory := conncrete_aggregated_transactions.CreateSignedTransactionsBuilderFactory(userSigBuilderFactory)
	aggTrsRepository := conncrete_aggregated_transactions.CreateSignedTransactionsRepository(objectRepository, aggregatedTrsRepository, aggregatedSignedTrsBuilderFactory)
	aggTrsService := conncrete_aggregated_transactions.CreateSignedTransactionsService(metaDataBuilderFactory, objBuilderFactory, objectService, aggregatedTrsService, storedTreeBuilderFactory)
	blkBuilderFactory := CreateBlockBuilderFactory(htBuilderFactory)

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateBlockRepository(blkBuilderFactory, htRepository, aggTrsRepository, objectRepository)
	service := CreateBlockService(storedTreesBuilderFactory, storedTreeBuilderFactory, metaDataBuilderFactory, objBuilderFactory, objectService, htService, aggTrsService)

	//make sure there is no transactions:
	_, noTrsErr := repository.Retrieve(basePath)
	if noTrsErr == nil {
		t.Errorf("there was supposed to be no block.")
	}

	//save the transaction:
	_, storedTrsErr := service.Save(basePath, blk)
	if storedTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedTrsErr.Error())
	}

	//retrieve the block:
	retBlk, retBlkErr := repository.Retrieve(basePath)
	if retBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retBlkErr.Error())
	}

	if !reflect.DeepEqual(blk, retBlk) {
		t.Errorf("the retrieved block is invalid")
	}

	//delete the transaction:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
