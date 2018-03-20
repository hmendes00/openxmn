package signed

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/users"
	convert "github.com/XMNBlockchain/exmachina-network/core/infrastructure/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateAtomicTransaction_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionsForTests()
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
	obj := CreateAtomicTransactionForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
