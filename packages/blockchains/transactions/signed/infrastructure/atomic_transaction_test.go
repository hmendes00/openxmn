package infrastructure

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	concrete_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateAtomicTransaction_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionsForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		[]byte(sig.GetSig().String()),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
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
	atomicTrs := createAtomicTransaction(met.(*concrete_metadata.MetaData), trs, sig)

	retMetaData := atomicTrs.GetMetaData()
	retTrs := atomicTrs.GetTransactions()
	retSig := atomicTrs.GetSignature()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned MetaData was invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned Transactions was invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned wallt signature was invalid")
	}

}

func TestCreateAtomicTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(AtomicTransaction)
	obj := CreateAtomicTransactionForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
