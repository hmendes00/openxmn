package infrastructure

import (
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	concrete_transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactionWithWalletSignature_Success(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	usrSig := concrete_users.CreateSignatureForTests(t)
	signedTrs := createTransaction(&id, trs, usrSig)

	retID := signedTrs.GetID()
	retTrs := signedTrs.GetTrs()
	retSig := signedTrs.GetSignature()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned id was invalid")
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
	obj := CreateTransactionForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
