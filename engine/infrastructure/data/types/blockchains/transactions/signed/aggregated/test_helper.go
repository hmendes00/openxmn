package aggregated

import (
	"strconv"
	"time"

	aggregated "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/transactions/signed/aggregated"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_stored_aggregated "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/blockchains/transactions/signed/aggregated"
	concrete_signed "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

// CreateTransactionsForTests creates an Transactions instance for tests
func CreateTransactionsForTests() *Transactions {
	//variables:
	id := uuid.NewV4()
	trs := concrete_signed.CreateTransactionsForTests()
	atomicTrs := concrete_signed.CreateAtomicTransactionsForTests()
	createdOn := time.Now().UTC()

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		atomicTrs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	aggregatedTrs := createTransactions(met.(*concrete_metadata.MetaData), trs, atomicTrs)
	return aggregatedTrs.(*Transactions)
}

// CreateSignedTransactionsForTests creates a SignedTransactions instance for tests
func CreateSignedTransactionsForTests() *SignedTransactions {
	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()
	cr := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

	sigTrs := createSignedTransactions(met.(*concrete_metadata.MetaData), trs, sig)
	return sigTrs.(*SignedTransactions)
}

// CreateTransactionsBuilderFactoryForTests creates a new TransactionsBuilderFactory for tests
func CreateTransactionsBuilderFactoryForTests() aggregated.TransactionsBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateTransactionsRepositoryForTests creates a new TransactionsRepository for tests
func CreateTransactionsRepositoryForTests() aggregated.TransactionsRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	signedTrsRepository := concrete_signed.CreateTransactionsRepositoryForTests()
	signedAtomicTrsRepository := concrete_signed.CreateAtomicTransactionsRepositoryForTests()
	aggregatedTrsBuilderFactory := CreateTransactionsBuilderFactoryForTests()
	out := CreateTransactionsRepository(metaDataRepository, signedTrsRepository, signedAtomicTrsRepository, aggregatedTrsBuilderFactory)
	return out
}

// CreateTransactionsServiceForTests creates a new TransactionsService for tests
func CreateTransactionsServiceForTests() aggregated.TransactionsService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	signedTrsService := concrete_signed.CreateTransactionsServiceForTests()
	atomicSignedTrsService := concrete_signed.CreateAtomicTransactionsServiceForTests()
	storedTrsBuilderFactory := concrete_stored_aggregated.CreateTransactionsBuilderFactory()
	out := CreateTransactionsService(metaDataService, signedTrsService, atomicSignedTrsService, storedTrsBuilderFactory)
	return out
}

// CreateSignedTransactionsBuilderFactoryForTests creates a new SignedTransactionsBuilderFactory for tests
func CreateSignedTransactionsBuilderFactoryForTests() aggregated.SignedTransactionsBuilderFactory {
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactoryForTests()
	out := CreateSignedTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	return out
}

// CreateSignedTransactionsRepositoryForTests creates a new SignedTransactionsRepository for tests
func CreateSignedTransactionsRepositoryForTests() aggregated.SignedTransactionsRepository {
	metaDataRepository := concrete_metadata.CreateMetaDataRepositoryForTests()
	userSigRepository := concrete_users.CreateSignatureRepositoryForTests()
	aggregatedTrsRepository := CreateTransactionsRepositoryForTests()
	signedTrsBuilderFactory := CreateSignedTransactionsBuilderFactoryForTests()
	out := CreateSignedTransactionsRepository(metaDataRepository, userSigRepository, aggregatedTrsRepository, signedTrsBuilderFactory)
	return out
}

// CreateSignedTransactionsServiceForTests creates a new SignedTransactionsService for tests
func CreateSignedTransactionsServiceForTests() aggregated.SignedTransactionsService {
	metaDataService := concrete_metadata.CreateMetaDataServiceForTests()
	userSigService := concrete_users.CreateSignatureServiceForTests()
	aggregatedTrsService := CreateTransactionsServiceForTests()
	storedAggregatedSignedTrsBuilderFactory := concrete_stored_aggregated.CreateSignedTransactionsBuilderFactory()
	out := CreateSignedTransactionsService(metaDataService, userSigService, aggregatedTrsService, storedAggregatedSignedTrsBuilderFactory)
	return out
}
