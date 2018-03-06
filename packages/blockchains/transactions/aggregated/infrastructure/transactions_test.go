package infrastructure

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactions_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := concrete_signed.CreateTransactionsForTests(t)
	atomicTrs := concrete_signed.CreateAtomicTransactionsForTests(t)
	createdOn := time.Now().UTC()

	htBlocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		atomicTrs.GetMetaData().GetHashTree().GetHash().Get(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(htBlocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()

	//execute:
	aggTrs := createTransactions(met.(*concrete_metadata.MetaData), trs, atomicTrs)

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

func TestCreateTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transactions)
	obj := CreateTransactionsForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
