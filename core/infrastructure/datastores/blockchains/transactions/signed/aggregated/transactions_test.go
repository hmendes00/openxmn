package aggregated

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
	concrete_signed "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/transactions/signed"
	convert "github.com/XMNBlockchain/exmachina-network/core/infrastructure/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactions_Success(t *testing.T) {

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
	obj := CreateTransactionsForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
