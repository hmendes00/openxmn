package infrastructure

import (
	"reflect"
	"testing"

	signed_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/lives/transactions/signed/infrastructure"
)

func TestCreateTransactionsBuilder_withAtomicTransactions_withTransactions_Success(t *testing.T) {

	//execute:
	trs := []signed_transactions.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	atomicTrs := []signed_transactions.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	build := createTransactionsBuilder()
	aggTrs, aggTrsErr := build.Create().WithAtomicTransactions(atomicTrs).WithTransactions(trs).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned")
	}

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

func TestCreateTransactionsBuilder_withAtomicTransactions_Success(t *testing.T) {

	//execute:
	atomicTrs := []signed_transactions.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	build := createTransactionsBuilder()
	aggTrs, aggTrsErr := build.Create().WithAtomicTransactions(atomicTrs).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned")
	}

	retTrs := aggTrs.GetTrs()
	retAtomicTrs := aggTrs.GetAtomicTrs()

	for index, oneAtomicTrs := range atomicTrs {
		if !reflect.DeepEqual(retAtomicTrs[index], oneAtomicTrs) {
			t.Errorf("the returned []AtomicTransaction was invalid at index: %d", index)
		}
	}

	if len(retTrs) > 0 {
		t.Errorf("the returned Transactions was expected to be empty.  Returned length: %d", len(retTrs))
	}
}

func TestCreateTransactionsBuilder_withTransactions_Success(t *testing.T) {

	//execute:
	trs := []signed_transactions.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	build := createTransactionsBuilder()
	aggTrs, aggTrsErr := build.Create().WithTransactions(trs).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned")
	}

	retTrs := aggTrs.GetTrs()
	retAtomicTrs := aggTrs.GetAtomicTrs()

	for index, oneTrs := range trs {
		if !reflect.DeepEqual(retTrs[index], oneTrs) {
			t.Errorf("the returned []Transaction was invalid at index: %d", index)
		}
	}

	if len(retAtomicTrs) > 0 {
		t.Errorf("the returned AtomicTransactions was expected to be empty.  Returned length: %d", len(retAtomicTrs))
	}

}
