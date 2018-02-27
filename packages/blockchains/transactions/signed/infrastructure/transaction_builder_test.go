package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_transactions "github.com/XMNBlockchain/core/packages/blockchains/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestBuildTransaction_Success(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if sigTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", sigTrsErr.Error())
	}

	retID := sigTrs.GetID()
	retTrs := sigTrs.GetTrs()
	retSig := sigTrs.GetSignature()
	retCreatedOn := sigTrs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID is invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned transaction is invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned user signature was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn time was invalid")
	}

}

func TestBuildTransaction_withoutID_returnsError(t *testing.T) {
	//variables:
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithTransaction(trs).WithSignature(sig).CreatedOn(createdOn).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}

func TestBuildTransaction_withoutTransaction_returnsError(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	sig := concrete_users.CreateSignatureForTests(t)
	createdOn := time.Now().UTC()

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).CreatedOn(createdOn).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}

func TestBuildTransaction_withoutSignature_returnsError(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	createdOn := time.Now().UTC()

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).CreatedOn(createdOn).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}

func TestBuildTransaction_withoutCreatedOn_returnsError(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).WithSignature(sig).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}