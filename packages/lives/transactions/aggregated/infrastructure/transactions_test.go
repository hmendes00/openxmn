package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	concrete_signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactions_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := []*concrete_signed.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	atomicTrs := []*concrete_signed.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	ht := concrete_hashtrees.CreateHashTreeForTests(t)

	//execute:
	aggTrs := createTransactions(&id, ht, trs, atomicTrs, createdOn)

	retID := aggTrs.GetID()
	retTrs := aggTrs.GetTrs()
	retAtomicTrs := aggTrs.GetAtomicTrs()
	retCreatedOn := aggTrs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID is invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	for index, oneTrs := range trs {
		if !reflect.DeepEqual(retTrs[index], oneTrs) {
			t.Errorf("the returned []Transaction was invalid at index: %d", index)
		}
	}

	for index, oneAtomicTrs := range atomicTrs {
		if !reflect.DeepEqual(retAtomicTrs[index], oneAtomicTrs) {
			t.Errorf("the returned []AtomicTransaction was invalid at index: %d", index)
		}
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned creation time is invalid")
	}

}

func TestCreateTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transactions)
	obj := CreateTransactionsForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
