package custom

import (
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateCustom_withCreate_Success(t *testing.T) {

	//variables:
	cr := CreateCreateForTests(t)

	//execute:
	cu := createCustomWithCreate(cr)
	retHasCreate := cu.HasCreate()
	retCreate := cu.GetCreate()

	if !retHasCreate {
		t.Errorf("the returned HasCustom has expected to be true, false returned")
	}

	if !reflect.DeepEqual(cr, retCreate) {
		t.Errorf("the returned create is invalid.")
	}

}

func TestCreateCustom_withCreate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Custom)
	obj := CreateCustomWithCreateForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
