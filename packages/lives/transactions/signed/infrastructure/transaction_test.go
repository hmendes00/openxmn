package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactionWithWalletSignature_Success(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	usrSig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()
	signedTrs := createTransaction(&id, trs, usrSig, createdOn)

	retID := signedTrs.GetID()
	retTrs := signedTrs.GetTrs()
	retSig := signedTrs.GetSignature()
	retCreatedOn := signedTrs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned id was invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned transactions was invalid")
	}

	if !reflect.DeepEqual(usrSig, retSig) {
		t.Errorf("the returned wallet signature was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn time was invalid")
	}

}

func TestCreateTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transaction)
	obj := CreateTransactionForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
