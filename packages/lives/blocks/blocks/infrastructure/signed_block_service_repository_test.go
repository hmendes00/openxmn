package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	concrete_chunks "github.com/XMNBlockchain/core/packages/lives/chunks/infrastructure"
	concrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/lives/metadata/infrastructure"
	concrete_aggregated_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/aggregated/infrastructure"
	concrete_signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	concrete_stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/infrastructure"
	concrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	concrete_stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/infrastructure"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/infrastructure"
	concrete_stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/infrastructure"
)

func TestSaveSignedBlock_thenRetrieve_Success(t *testing.T) {

	//create the block:
	signedBlk := CreateSignedBlockForTests(t)

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
	trsRepository := concrete_transactions.CreateTransactionRepository(chkRepository)
	storedTrsBuilderFactory := concrete_stored_transactions.CreateTransactionBuilderFactory()
	trsService := concrete_transactions.CreateTransactionService(metaDataBuilderFactory, metaDataService, chksBuilderFactory, chkService, storedTrsBuilderFactory)
	signedTrsBuilderFactory := concrete_signed_transactions.CreateTransactionBuilderFactory()
	atomicTrsBuilderFactory := concrete_signed_transactions.CreateAtomicTransactionBuilderFactory(htBuilderFactory)
	userSigRepository := concrete_users.CreateSignatureRepository(fileRepository)
	storedSignedTrsBuilderFactory := concrete_stored_signed_transactions.CreateTransactionBuilderFactory()
	userSigService := concrete_users.CreateSignatureService(fileService, fileBuilderFactory)
	publicKeyBuilderFactory := concrete_cryptography.CreatePublicKeyBuilderFactory()
	sigBuilderFactory := concrete_cryptography.CreateSignatureBuilderFactory(publicKeyBuilderFactory)
	userBuilderFactory := concrete_users.CreateUserBuilderFactory()
	userSigBuilderFactory := concrete_users.CreateSignatureBuilderFactory(sigBuilderFactory, userBuilderFactory)
	signedTrsRepository := concrete_signed_transactions.CreateTransactionRepository(metaDataRepository, userSigRepository, trsRepository, signedTrsBuilderFactory)
	signedTrsService := concrete_signed_transactions.CreateTransactionService(metaDataBuilderFactory, metaDataService, trsService, storedSignedTrsBuilderFactory, userSigService)
	atomicTrsRepository := concrete_signed_transactions.CreateAtomicTransactionRepository(metaDataRepository, userSigRepository, htRepository, trsRepository, atomicTrsBuilderFactory)
	storedAtomicTrsBuilderFactory := concrete_stored_signed_transactions.CreateAtomicTransactionBuilderFactory()
	atomicTrsService := concrete_signed_transactions.CreateAtomicTransactionService(metaDataBuilderFactory, metaDataService, htService, userSigService, trsService, storedAtomicTrsBuilderFactory)
	aggregatedTrsBuilderFactory := concrete_aggregated_transactions.CreateTransactionsBuilderFactory(htBuilderFactory)
	storedAggregatedTrsBuilderFactory := concrete_stored_aggregated_transactions.CreateTransactionsBuilderFactory()
	aggregatedTrsRepository := concrete_aggregated_transactions.CreateTransactionsRepository(metaDataRepository, htRepository, signedTrsRepository, atomicTrsRepository, aggregatedTrsBuilderFactory)
	aggregatedSignedTrsBuilderFactory := concrete_aggregated_transactions.CreateSignedTransactionsBuilderFactory(userSigBuilderFactory)
	aggregatedTrsService := concrete_aggregated_transactions.CreateTransactionsService(metaDataBuilderFactory, metaDataService, htService, signedTrsService, atomicTrsService, storedAggregatedTrsBuilderFactory)
	storedAggregatedSignedTrsBuilderFactory := concrete_stored_aggregated_transactions.CreateSignedTransactionsBuilderFactory()
	aggrSignedTrsRepository := concrete_aggregated_transactions.CreateSignedTransactionsRepository(metaDataRepository, userSigRepository, aggregatedTrsRepository, aggregatedSignedTrsBuilderFactory)
	aggrSignedTrsService := concrete_aggregated_transactions.CreateSignedTransactionsService(metaDataBuilderFactory, metaDataService, userSigService, aggregatedTrsService, storedAggregatedSignedTrsBuilderFactory)
	storedBlkBuilderFactory := concrete_stored_blocks.CreateBlockBuilderFactory()
	blkBuilderFactory := CreateBlockBuilderFactory(htBuilderFactory)
	blkRepository := CreateBlockRepository(metaDataRepository, htRepository, aggrSignedTrsRepository, blkBuilderFactory)
	blkService := CreateBlockService(metaDataBuilderFactory, metaDataService, htService, aggrSignedTrsService, storedBlkBuilderFactory)
	signedBlkBuilderFactory := CreateSignedBlockBuilderFactory()
	storedSignedBlkBuilderFactory := concrete_stored_blocks.CreateSignedBlockBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateSignedBlockRepository(metaDataRepository, userSigRepository, blkRepository, signedBlkBuilderFactory)
	service := CreateSignedBlockService(metaDataBuilderFactory, metaDataService, userSigService, blkService, storedSignedBlkBuilderFactory)

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
