package infrastructure

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactionsBuilder_withAtomicTransactions_withTransactions_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := concrete_signed.CreateTransactionsForTests(t)
	atomicTrs := concrete_signed.CreateAtomicTransactionsForTests(t)

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		atomicTrs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	build := createTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	aggTrs, aggTrsErr := build.Create().WithID(&id).WithAtomicTransactions(atomicTrs).WithTransactions(trs).CreatedOn(createdOn).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned %s", aggTrsErr.Error())
	}

	retMetaData := aggTrs.GetMetaData()
	retTrs := aggTrs.GetTransactions()
	retAtomicTrs := aggTrs.GetAtomicTransactions()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned ID is invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned signed transactions is invalid")
	}

	if !reflect.DeepEqual(atomicTrs, retAtomicTrs) {
		t.Errorf("the returned signed atomic transactions is invalid")
	}
}

func TestCreateTransactionsBuilder_withAtomicTransactions_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	atomicTrs := concrete_signed.CreateAtomicTransactionsForTests(t)

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		atomicTrs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	build := createTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	aggTrs, aggTrsErr := build.Create().WithID(&id).WithAtomicTransactions(atomicTrs).CreatedOn(createdOn).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned %s", aggTrsErr.Error())
	}

	retMetaData := aggTrs.GetMetaData()
	retHasTransactions := aggTrs.HasTransactions()
	retAtomicTrs := aggTrs.GetAtomicTransactions()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned ID is invalid")
	}

	if retHasTransactions {
		t.Errorf("there should be no transactions")
	}

	if !reflect.DeepEqual(atomicTrs, retAtomicTrs) {
		t.Errorf("the returned signed atomic transactions is invalid")
	}
}

func TestCreateTransactionsBuilder_withTransactions_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := concrete_signed.CreateTransactionsForTests(t)

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	build := createTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	aggTrs, aggTrsErr := build.Create().WithID(&id).WithTransactions(trs).CreatedOn(createdOn).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned %s", aggTrsErr.Error())
	}

	retMetaData := aggTrs.GetMetaData()
	retTrs := aggTrs.GetTransactions()
	retHasAtomicTrs := aggTrs.HasAtomicTransactions()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned ID is invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned signed transactions is invalid")
	}

	if retHasAtomicTrs {
		t.Errorf("there should be no atomic transactions")
	}

}
