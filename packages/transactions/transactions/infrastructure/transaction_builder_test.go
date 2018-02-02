package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_body "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body"
	uuid "github.com/satori/go.uuid"
)

func TestCreateBuilder_withUUID_withBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()
	bod := concrete_body.CreateBodyWithCustomForTests(t)

	//execute:
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithID(&id).WithBody(bod).CreatedOn(createdOn).Now()

	if trsErr != nil {
		t.Errorf("the returned error was expected to be nil, Returned: %s", trsErr.Error())
	}

	if trs == nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

	retID := trs.GetID()
	retBody := trs.GetBody()
	retCreatedOn := trs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned id was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(bod, retBody) {
		t.Errorf("the returned body was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn was invalid")
	}

}

func TestCreateBuilder_withoutUUID_withBody_withoutCreatedOn_Success(t *testing.T) {

	//execute:
	bod := concrete_body.CreateBodyWithCustomForTests(t)
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithBody(bod).Now()

	if trsErr != nil {
		t.Errorf("the returned error was expected to be nil, Returned: %s", trsErr.Error())
	}

	if trs == nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

	retBody := trs.GetBody()

	if !reflect.DeepEqual(bod, retBody) {
		t.Errorf("the returned body was invalid")
	}

}

func TestCreateBuilder_withUUID_withoutBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()

	//execute:
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithID(&id).CreatedOn(createdOn).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}
}
