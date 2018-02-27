package infrastructure

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreate_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()
	obj := JsDataForTests{
		Name:        "Some name",
		Description: "This is some description",
	}

	js, _ := json.Marshal(&obj)

	//execute:
	trs := createTransaction(&id, js, createdOn)
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

func TestCreateTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transaction)
	obj := CreateTransactionForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
