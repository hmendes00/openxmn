package infrastructure

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	concrete_body "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body"
	uuid "github.com/satori/go.uuid"
)

func TestCreateBuilder_withUUID_withKarma_withBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()
	karma := rand.Int() % 20
	bod := concrete_body.CreateBodyWithCustomForTests(t)

	//execute:
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithID(&id).WithBody(bod).WithKarma(karma).CreatedOn(createdOn).Now()

	if trsErr != nil {
		t.Errorf("the returned error was expected to be nil, Returned: %s", trsErr.Error())
	}

	if trs == nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

	retID := trs.GetID()
	retKarma := trs.GetKarma()
	retBody := trs.GetBody()
	retCreatedOn := trs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned id was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(bod, retBody) {
		t.Errorf("the returned body was invalid")
	}

	if !reflect.DeepEqual(karma, retKarma) {
		t.Errorf("the returned karma was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn was invalid")
	}

}

func TestCreateBuilder_withoutUUID_withKarma_withBody_withCreatedOn_Success(t *testing.T) {

	//execute:
	createdOn := time.Now()
	karma := rand.Int() % 20
	bod := concrete_body.CreateBodyWithCustomForTests(t)
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithBody(bod).WithKarma(karma).CreatedOn(createdOn).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

}

func TestCreateBuilder_withUUID_withKarma_withBody_withoutCreatedOn_Success(t *testing.T) {

	//execute:
	id := uuid.NewV4()
	karma := rand.Int() % 20
	bod := concrete_body.CreateBodyWithCustomForTests(t)
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithID(&id).WithKarma(karma).WithBody(bod).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

}

func TestCreateBuilder_withUUID_withoutBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	karma := rand.Int() % 20
	createdOn := time.Now()

	//execute:
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithKarma(karma).WithID(&id).CreatedOn(createdOn).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}
}
