package aggregated

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateSignedTransactions_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests()
	sig := users.CreateSignatureForTests()
	cr := time.Now()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		sig.GetMetaData().GetID().Bytes(),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

	//execute:
	sigTrs := createSignedTransactions(met.(*concrete_metadata.MetaData), trs, sig)

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
		t.Errorf("the returned pointer Signature is invalid")
	}

}

func TestCreateSignedTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(SignedTransactions)
	obj := CreateSignedTransactionsForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
