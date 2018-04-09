package aggregated

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	uuid "github.com/satori/go.uuid"
)

func TestBuildSignedTransactions_withID_withSignature_withTransactions_createdOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()
	cr := time.Now().UTC()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	build := createSignedTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).WithTransactions(trs).CreatedOn(cr).Now()

	if sigTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", sigTrsErr.Error())
	}

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetID().Bytes(),
	}

	ht, _ := htBuilderFactory.Create().Create().WithBlocks(blocks).Now()
	met, _ := metaDataBuilderFactory.Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

	retMetaData := sigTrs.GetMetaData()
	retTrs := sigTrs.GetTransactions()
	retSig := sigTrs.GetSignature()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned MetaData is invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned Transactions is invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned user Signature is invalid")
	}

}

func TestBuildSignedTransactions_withoutCreatedOn_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	build := createSignedTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).WithTransactions(trs).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildSignedTransactions_withoutID_returnsError(t *testing.T) {

	//variables:
	trs := CreateTransactionsForTests()
	sig := concrete_users.CreateSignatureForTests()
	cr := time.Now()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	build := createSignedTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithSignature(sig).WithTransactions(trs).CreatedOn(cr).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildSignedTransactions_withoutSignature_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests()
	cr := time.Now()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	build := createSignedTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransactions(trs).CreatedOn(cr).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildSignedTransactions_withoutTransactions_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	sig := concrete_users.CreateSignatureForTests()
	cr := time.Now()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateBuilderFactory()
	build := createSignedTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).CreatedOn(cr).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}
