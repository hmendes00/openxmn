package signed

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
	concrete_transactions "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/transactions"
	concrete_users "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/users"
	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactionWithWalletSignature_Success(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests()
	usrSig := concrete_users.CreateSignatureForTests()
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
		trs.GetMetaData().GetHashTree().GetHash().Get(),
		usrSig.GetMetaData().GetHashTree().GetHash().Get(),
	}
	ht, htErr := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()
	if htErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", htErr.Error())
	}

	met, metErr := concrete_metadata.CreateMetaDataBuilderFactory().Create().Create().WithID(&id).WithHashTree(ht).CreatedOn(createdOn).Now()
	if metErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", metErr.Error())
	}

	signedTrs := createTransaction(met.(*concrete_metadata.MetaData), trs, usrSig)

	retMetaData := signedTrs.GetMetaData()
	retTrs := signedTrs.GetTransaction()
	retSig := signedTrs.GetSignature()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned metadata was invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned transactions was invalid")
	}

	if !reflect.DeepEqual(usrSig, retSig) {
		t.Errorf("the returned wallet signature was invalid")
	}

}

func TestCreateTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transaction)
	obj := CreateTransactionForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
