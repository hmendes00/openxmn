package custom

import (
	"encoding/json"
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateCreate_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	insID := uuid.NewV4()
	ins := InstanceForTests{
		ID:          &insID,
		Name:        "some name",
		Description: "this is some custom description",
	}
	js, _ := json.Marshal(ins)

	//execute:
	cr := createCreate(&id, js)
	retID := cr.GetID()
	retJS := cr.GetJSON()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID is invalid.  Returned: %s, Expected: %s", retID.String(), id.String())
	}

	if !reflect.DeepEqual(js, retJS) {
		t.Errorf("the returned JSON is invalid.")
	}

}

func TestCreateCreate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Create)
	obj := CreateCreateForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
