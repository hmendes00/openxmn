package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/domain"
	concrete_transactions "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestBuildAtomicTransaction_Success(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := []transactions.Transaction{
		concrete_transactions.CreateTransactionForTests(t),
		concrete_transactions.CreateTransactionForTests(t),
	}

	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithTransactions(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if atomicTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", atomicTrsErr.Error())
	}

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
		t.Errorf("the returned user signature was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn time was invalid")
	}

}

func TestBuildAtomicTransaction_withoutID_returnsError(t *testing.T) {

	//execute:
	trs := []transactions.Transaction{
		concrete_transactions.CreateTransactionForTests(t),
		concrete_transactions.CreateTransactionForTests(t),
	}

	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithTransactions(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}

}

func TestBuildAtomicTransaction_withoutTransactions_returnsError(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithSignature(sig).CreatedOn(createdOn).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}

}

func TestBuildAtomicTransaction_withoutSignature_returnsError(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := []transactions.Transaction{
		concrete_transactions.CreateTransactionForTests(t),
		concrete_transactions.CreateTransactionForTests(t),
	}

	createdOn := time.Now().UTC()

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithTransactions(trs).CreatedOn(createdOn).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}

}

func TestBuildAtomicTransaction_withoutCreatedOn_returnsError(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	trs := []transactions.Transaction{
		concrete_transactions.CreateTransactionForTests(t),
		concrete_transactions.CreateTransactionForTests(t),
	}

	sig := concrete_users.CreateSignatureForTests(t)

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()

	build := createAtomicTransactionBuilder(htBuilderFactory)
	atomicTrs, atomicTrsErr := build.Create().WithID(&id).WithTransactions(trs).WithSignature(sig).Now()

	if atomicTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if atomicTrs != nil {
		t.Errorf("the returned atomicTrs was expected to be nil, instance returned")
	}
}
