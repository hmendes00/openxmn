package aggregated

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_signed "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions/signed"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactionsBuilder_withAtomicTransactions_withTransactions_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := concrete_signed.CreateTransactionsForTests()
	atomicTrs := concrete_signed.CreateAtomicTransactionsForTests()

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		atomicTrs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

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
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	atomicTrs := concrete_signed.CreateAtomicTransactionsForTests()

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		atomicTrs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

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
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := concrete_signed.CreateTransactionsForTests()

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

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
