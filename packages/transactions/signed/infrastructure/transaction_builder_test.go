package infrastructure

import (
	"reflect"
	"testing"

	concrete_transactions "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure"
	concrete_users "github.com/XMNBlockchain/core/packages/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestBuildTransaction_Success(t *testing.T) {
	//variables:
	id := uuid.NewV4()
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).WithSignature(sig).Now()

	if sigTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", sigTrsErr.Error())
	}

	retID := sigTrs.GetID()
	retTrs := sigTrs.GetTrs()
	retSig := sigTrs.GetSignature()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID is invalid")
	}

	if !reflect.DeepEqual(trs, retTrs) {
		t.Errorf("the returned transaction is invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned user signature was invalid")
	}

}

func TestBuildTransaction_withoutID_returnsError(t *testing.T) {
	//variables:
	trs := concrete_transactions.CreateTransactionForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithTransaction(trs).WithSignature(sig).Now()

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

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).Now()

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

	//execute:
	build := createTransactionBuilder()
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransaction(trs).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned transactiom was expected to be nil, instance returned")
	}

}
