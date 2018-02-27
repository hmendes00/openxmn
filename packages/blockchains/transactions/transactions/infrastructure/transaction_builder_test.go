package infrastructure

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	uuid "github.com/satori/go.uuid"
)

func TestCreateBuilder_withUUID_withBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	//execute:
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithID(&id).WithJSON(js).CreatedOn(createdOn).Now()

	if trsErr != nil {
		t.Errorf("the returned error was expected to be nil, Returned: %s", trsErr.Error())
	}

	if trs == nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

	retID := trs.GetID()
	retJS := trs.GetJSON()
	retCreatedOn := trs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned id was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(js, retJS) {
		t.Errorf("the returned json was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn was invalid")
	}

}

func TestCreateBuilder_withoutUUID_withBody_withCreatedOn_Success(t *testing.T) {

	//variables:
	createdOn := time.Now()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	//execute:
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithJSON(js).CreatedOn(createdOn).Now()

	if trsErr == nil {
		t.Errorf("the error was expected to be an error, nil returned")
	}

	if trs != nil {
		t.Errorf("the returned transaction was expected to be an instance, nil returned")
	}

}

func TestCreateBuilder_withUUID_withBody_withoutCreatedOn_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	//execute:
	build := createTransactionBuilder()
	trs, trsErr := build.Create().WithID(&id).WithJSON(js).Now()

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
