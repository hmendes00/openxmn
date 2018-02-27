package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	signed_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/domain"
	concrete_signed "github.com/XMNBlockchain/core/packages/blockchains/transactions/signed/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestCreateTransactionsBuilder_withAtomicTransactions_withTransactions_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := []signed_transactions.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	atomicTrs := []signed_transactions.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	build := createTransactionsBuilder(htBuilderFactory)
	aggTrs, aggTrsErr := build.Create().WithID(&id).WithAtomicTransactions(atomicTrs).WithTransactions(trs).CreatedOn(createdOn).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned %s", aggTrsErr.Error())
	}

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

func TestCreateTransactionsBuilder_withAtomicTransactions_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	atomicTrs := []signed_transactions.AtomicTransaction{
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
		concrete_signed.CreateAtomicTransactionForTests(t),
	}

	build := createTransactionsBuilder(htBuilderFactory)
	aggTrs, aggTrsErr := build.Create().WithID(&id).WithAtomicTransactions(atomicTrs).CreatedOn(createdOn).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned %s", aggTrsErr.Error())
	}

	retID := aggTrs.GetID()
	retTrs := aggTrs.GetTrs()
	retAtomicTrs := aggTrs.GetAtomicTrs()
	retCreatedOn := aggTrs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID is invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	for index, oneAtomicTrs := range atomicTrs {
		if !reflect.DeepEqual(retAtomicTrs[index], oneAtomicTrs) {
			t.Errorf("the returned []AtomicTransaction was invalid at index: %d", index)
		}
	}

	if len(retTrs) > 0 {
		t.Errorf("the returned Transactions was expected to be empty.  Returned length: %d", len(retTrs))
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned creation time is invalid")
	}
}

func TestCreateTransactionsBuilder_withTransactions_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	//execute:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()
	trs := []signed_transactions.Transaction{
		concrete_signed.CreateTransactionForTests(t),
		concrete_signed.CreateTransactionForTests(t),
	}

	build := createTransactionsBuilder(htBuilderFactory)
	aggTrs, aggTrsErr := build.Create().WithID(&id).WithTransactions(trs).CreatedOn(createdOn).Now()
	if aggTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", aggTrsErr.Error())
	}

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

	if len(retAtomicTrs) > 0 {
		t.Errorf("the returned AtomicTransactions was expected to be empty.  Returned length: %d", len(retAtomicTrs))
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned creation time is invalid")
	}

}
