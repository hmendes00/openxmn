package infrastructure

import (
	"reflect"
	"testing"
	"time"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	concrete_transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestCreateAtomicTransaction_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	trs := []*concrete_transactions.Transaction{
		concrete_transactions.CreateTransactionForTests(t),
		concrete_transactions.CreateTransactionForTests(t),
	}
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	//execute:
	atomicTrs := createAtomicTransaction(&id, trs, sig, createdOn)

	retID := atomicTrs.GetID()
	retTrs := atomicTrs.GetTrs()
	retSig := atomicTrs.GetSignature()
	retCreatedOn := atomicTrs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid")
	}

	for index, oneTrs := range trs {
		if !reflect.DeepEqual(retTrs[index], oneTrs) {
			t.Errorf("the returned []transaction was invalid at index: %d", index)
		}
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned wallt signature was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn timestamp was invalid")
	}

}

func TestCreateAtomicTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(AtomicTransaction)
	obj := CreateAtomicTransactionForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
