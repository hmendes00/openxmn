package signed

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

func TestBuildAtomicTransaction_Success(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory, metaDataBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithTransactions(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if atomicTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", atomicTrsErr.Error())
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, htErr := htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
	if htErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", htErr.Error())
	}

	met, metErr := metaDataBuilderFactory.Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()
	if metErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", metErr.Error())
	}

	retMetaData := atomicTrs.GetMetaData()
	retTrs := atomicTrs.GetTransactions()
	retSig := atomicTrs.GetSignature()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned metadata was invalid")
	}

	if !reflect.DeepEqual(retTrs, trs) {
		t.Errorf("the returned Transactions was invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned user signature was invalid")
	}

}

func TestBuildAtomicTransaction_withoutID_returnsError(t *testing.T) {

	//execute:
	trs := concrete_transactions.CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory, metaDataBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithTransactions(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}

}

func TestBuildAtomicTransaction_withoutTransactions_returnsError(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory, metaDataBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithSignature(sig).CreatedOn(createdOn).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}

}

func TestBuildAtomicTransaction_withoutSignature_returnsError(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionsForTests()
	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory, metaDataBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithTransactions(trs).CreatedOn(createdOn).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}

}

func TestBuildAtomicTransaction_withoutCreatedOn_returnsError(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory, metaDataBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithTransactions(trs).WithSignature(sig).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}
}
