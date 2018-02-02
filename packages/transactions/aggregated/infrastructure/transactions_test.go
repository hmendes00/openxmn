package infrastructure

import (
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	concrete_signed "github.com/XMNBlockchain/core/packages/transactions/signed/infrastructure"
)

func TestCreateTransactions_Success(t *testing.T) {

	//variables:
	trs := []*concrete_signed.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	atomicTrs := []*concrete_signed.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	//execute:
	aggTrs := createTransactions(trs, atomicTrs)

	retTrs := aggTrs.GetTrs()
	retAtomicTrs := aggTrs.GetAtomicTrs()

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

}

func TestCreateTransactions_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transactions)
	obj := CreateTransactionsForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
