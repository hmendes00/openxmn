package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_users "github.com/XMNBlockchain/core/packages/lives/users/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestBuildSignedTransactions_withID_withSignature_withTransactions_createdOn_Success(t *testing.T) {

	//variables:
	usrSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	id := uuid.NewV4()
	trs := CreateTransactionsForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	cr := time.Now()

	//execute:
	build := createSignedTransactionsBuilder(usrSigBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).WithTransactions(trs).CreatedOn(cr).Now()

	if sigTrsErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", sigTrsErr.Error())
	}

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
		t.Errorf("the returned user Signature is invalid")
	}

	if !reflect.DeepEqual(cr, retCr) {
		t.Errorf("the returned createdOn time is invalid")
	}

}

func TestBuildSignedTransactions_withoutCreatedOn_returnsError(t *testing.T) {

	//variables:
	usrSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	id := uuid.NewV4()
	trs := CreateTransactionsForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)

	//execute:
	build := createSignedTransactionsBuilder(usrSigBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).WithTransactions(trs).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildSignedTransactions_withoutID_returnsError(t *testing.T) {

	//variables:
	usrSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	trs := CreateTransactionsForTests(t)
	sig := concrete_users.CreateSignatureForTests(t)
	cr := time.Now()

	//execute:
	build := createSignedTransactionsBuilder(usrSigBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithSignature(sig).WithTransactions(trs).CreatedOn(cr).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildSignedTransactions_withoutSignature_returnsError(t *testing.T) {

	//variables:
	usrSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	id := uuid.NewV4()
	trs := CreateTransactionsForTests(t)
	cr := time.Now()

	//execute:
	build := createSignedTransactionsBuilder(usrSigBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithTransactions(trs).CreatedOn(cr).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildSignedTransactions_withoutTransactions_returnsError(t *testing.T) {

	//variables:
	usrSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	id := uuid.NewV4()
	sig := concrete_users.CreateSignatureForTests(t)
	cr := time.Now()

	//execute:
	build := createSignedTransactionsBuilder(usrSigBuilderFactory)
	sigTrs, sigTrsErr := build.Create().WithID(&id).WithSignature(sig).CreatedOn(cr).Now()

	if sigTrsErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if sigTrs != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}
