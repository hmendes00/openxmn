package infrastructure

import (
	"reflect"
	"testing"
	"time"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestCreateSignedTransactions_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := CreateTransactionsForTests(t)
	sig := users.CreateSignatureForTests(t)
	cr := time.Now()

	//execute:
	sigTrs := createSignedTransactions(&id, trs, sig, cr)

	retID := sigTrs.GetID()
	retTrs := sigTrs.GetTrs()
	retSig := sigTrs.GetSignature()
	retCr := sigTrs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID is invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned Transactions is invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned pointer Signature is invalid")
	}

	if !reflect.DeepEqual(cr, retCr) {
		t.Errorf("the returned createdOn time is invalid")
	}

}

func TestCreateSignedTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(SignedTransactions)
	obj := CreateSignedTransactionsForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
