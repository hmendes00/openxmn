package infrastructure

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
	users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateSignedTransactions_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests(t)
	sig := users.CreateSignatureForTests(t)
	cr := time.Now()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(cr.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		[]byte(sig.GetSig().String()),
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	met, _ := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(cr).Now()

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
	obj := CreateSignedTransactionsForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
