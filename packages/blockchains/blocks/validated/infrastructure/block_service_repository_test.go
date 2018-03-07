package infrastructure

import (
	"path/filepath"
	"reflect"
	"testing"

	concrete_blocks "github.com/XMNBlockchain/core/packages/blockchains/blocks/blocks/infrastructure"
	concrete_chunks "github.com/XMNBlockchain/core/packages/blockchains/chunks/infrastructure"
	concrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_aggregated "github.com/XMNBlockchain/core/packages/blockchains/transactions/aggregated/infrastructure"
	concrete_signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	concrete_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	concrete_stored_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/blocks/infrastructure"
	concrete_stored_validated_blocks "github.com/XMNBlockchain/core/packages/storages/blocks/validated/infrastructure"
	concrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	concrete_stored_aggregated_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/aggregated/infrastructure"
	concrete_stored_signed_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/signed/infrastructure"
	concrete_stored_transactions "github.com/XMNBlockchain/core/packages/storages/transactions/transactions/infrastructure"
	concrete_stored_users "github.com/XMNBlockchain/core/packages/storages/users/infrastructure"
)

func TestSaveValidatedBlock_thenRetrieve_Success(t *testing.T) {

	//create the block:
	validatedBlk := CreateBlockForTests(t)

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
	transBuilderFactory := concrete_transactions.CreateTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	trsRepository := concrete_transactions.CreateTransactionRepository(chkRepository, metaDataRepository, trsBuilderFactory)
	storedTrsBuilderFactory := concrete_stored_transactions.CreateTransactionBuilderFactory()
	storedTrsansBuilderFactory := concrete_stored_transactions.CreateTransactionsBuilderFactory()
	trsService := concrete_transactions.CreateTransactionService(metaDataService, chksBuilderFactory, chkService, storedTrsBuilderFactory)
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
	userSigsBuilderFactory := concrete_users.CreateSignaturesBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	userSigsRepository := concrete_users.CreateSignaturesRepository(metaDataRepository, userSigRepository, userSigsBuilderFactory)
	storedSigsBuilderFactory := concrete_stored_users.CreateSignaturesBuilderFactory()
	userSigsService := concrete_users.CreateSignaturesService(metaDataService, userSigService, storedSigsBuilderFactory)
	transRepository := concrete_transactions.CreateTransactionsRepository(metaDataRepository, trsRepository, transBuilderFactory)
	transService := concrete_transactions.CreateTransactionsService(metaDataService, trsService, storedTrsansBuilderFactory)
	atomicTrsBuilderFactory := concrete_signed_transactions.CreateAtomicTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedAtomicTrsBuilderFactory := concrete_stored_signed_transactions.CreateAtomicTransactionBuilderFactory()
	atomicTrsRepository := concrete_signed_transactions.CreateAtomicTransactionRepository(metaDataRepository, userSigRepository, transRepository, atomicTrsBuilderFactory)
	atomicTrsService := concrete_signed_transactions.CreateAtomicTransactionService(metaDataService, userSigService, transService, storedAtomicTrsBuilderFactory)
	atomicTransBuilderFactory := concrete_signed_transactions.CreateAtomicTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedAtomicTransBuilderFactory := concrete_stored_signed_transactions.CreateAtomicTransactionsBuilderFactory()
	atomicTransRepository := concrete_signed_transactions.CreateAtomicTransactionsRepository(metaDataRepository, atomicTrsRepository, atomicTransBuilderFactory)
	signedTrsBuilderFactory := concrete_signed_transactions.CreateTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	transactionRepository := concrete_signed_transactions.CreateTransactionRepository(metaDataRepository, userSigRepository, trsRepository, signedTrsBuilderFactory)
	storedSignedTrsBuilderFactory := concrete_stored_signed_transactions.CreateTransactionBuilderFactory()
	transactionService := concrete_signed_transactions.CreateTransactionService(metaDataService, trsService, storedSignedTrsBuilderFactory, userSigService)
	signedTransBuilderFactory := concrete_signed_transactions.CreateTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	signedStoredTransBuilderFactory := concrete_stored_signed_transactions.CreateTransactionsBuilderFactory()
	signedTransRepository := concrete_signed_transactions.CreateTransactionsRepository(metaDataRepository, transactionRepository, signedTransBuilderFactory)
	signedTransService := concrete_signed_transactions.CreateTransactionsService(metaDataService, transactionService, signedStoredTransBuilderFactory)
	atomicTransService := concrete_signed_transactions.CreateAtomicTransactionsService(metaDataService, atomicTrsService, storedAtomicTransBuilderFactory)
	aggregatedTrsBuilderFactory := concrete_aggregated.CreateTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	aggregatedStoredTrsBuilderFactory := concrete_stored_aggregated_transactions.CreateTransactionsBuilderFactory()
	aggregatedTrsRepository := concrete_aggregated.CreateTransactionsRepository(metaDataRepository, signedTransRepository, atomicTransRepository, aggregatedTrsBuilderFactory)
	aggregatedTrsService := concrete_aggregated.CreateTransactionsService(metaDataService, signedTransService, atomicTransService, aggregatedStoredTrsBuilderFactory)
	aggregatedSignedTrsBuilderFactory := concrete_aggregated.CreateSignedTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedAggregatedSignedTrsBuilderFactory := concrete_stored_aggregated_transactions.CreateSignedTransactionsBuilderFactory()
	aggrSignedTrsRepository := concrete_aggregated.CreateSignedTransactionsRepository(metaDataRepository, userSigRepository, aggregatedTrsRepository, aggregatedSignedTrsBuilderFactory)
	aggrSignedTrsService := concrete_aggregated.CreateSignedTransactionsService(metaDataService, userSigService, aggregatedTrsService, storedAggregatedSignedTrsBuilderFactory)
	blkBuilderFactory := concrete_blocks.CreateBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	blkRepository := concrete_blocks.CreateBlockRepository(metaDataRepository, aggrSignedTrsRepository, blkBuilderFactory)
	storedBlkBuilderFactory := concrete_stored_blocks.CreateBlockBuilderFactory()
	blkService := concrete_blocks.CreateBlockService(metaDataService, aggrSignedTrsService, storedBlkBuilderFactory)
	signedBlkBuilderFactory := concrete_blocks.CreateSignedBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedSignedBlkBuilderFactory := concrete_stored_blocks.CreateSignedBlockBuilderFactory()
	signedBlkRepository := concrete_blocks.CreateSignedBlockRepository(metaDataRepository, userSigRepository, blkRepository, signedBlkBuilderFactory)
	signedBlkService := concrete_blocks.CreateSignedBlockService(metaDataService, userSigService, blkService, storedSignedBlkBuilderFactory)
	validatedBlkBuilderFactory := CreateBlockBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	storedValidatedBlkBuilderFactory := concrete_stored_validated_blocks.CreateBlockBuilderFactory()

	//delete the files folder at the end:
	defer func() {
		fileService.DeleteAll(basePath)
	}()

	//execute:
	repository := CreateBlockRepository(metaDataRepository, signedBlkRepository, userSigsRepository, validatedBlkBuilderFactory)
	service := CreateBlockService(metaDataService, signedBlkService, userSigsService, storedValidatedBlkBuilderFactory)

	//make sure there is no blocks:
	_, noTrsErr := repository.Retrieve(basePath)
	if noTrsErr == nil {
		t.Errorf("there was supposed to be no signed block.")
	}

	//save the block:
	_, storedTrsErr := service.Save(basePath, validatedBlk)
	if storedTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedTrsErr.Error())
	}

	//retrieve the block:
	retBlk, retBlkErr := repository.Retrieve(basePath)
	if retBlkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retBlkErr.Error())
	}

	if !reflect.DeepEqual(validatedBlk, retBlk) {
		t.Errorf("the retrieved block is invalid")
	}

	//delete the block:
	delErr := fileService.DeleteAll(basePath)
	if delErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
	}
}
