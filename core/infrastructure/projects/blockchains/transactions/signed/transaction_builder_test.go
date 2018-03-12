package signed

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/users"
	uuid "github.com/satori/go.uuid"
)

func TestBuildTransaction_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetHashTree().GetHash().Get(),
	}
	ht, htErr := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	if htErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", htErr.Error())
	}

	met, metErr := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()
	if metErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", metErr.Error())
	}

	//execute:
	build := createTransactionBuilder(htBuilderFactory, metaBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if sigTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", sigTrsErr.Error())
	}

	retMetaData := sigTrs.GetMetaData()
	retTrs := sigTrs.GetTransaction()
	retSig := sigTrs.GetSignature()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned metadata is invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned transaction is invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned user signature was invalid")
	}

}

func TestBuildTransaction_withoutID_returnsError(t *testing.T) {
	//variables:
	trs := concrete_transactions.CreateTransactionForTests()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metaBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithTransaction(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}

func TestBuildTransaction_withoutTransaction_returnsError(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	sig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metaBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).CreatedOn(createdOn).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}

func TestBuildTransaction_withoutSignature_returnsError(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests()
	createdOn := time.Now().UTC()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metaBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).CreatedOn(createdOn).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}

func TestBuildTransaction_withoutCreatedOn_returnsError(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests()
	sig := concrete_users.CreateSignatureForTests()

	//execute:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metaBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).WithSignature(sig).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}
