package servers

import (
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateServer_withCreate_Success(t *testing.T) {

	//variables:
	cr := CreateCreateForTests(t)

	//execute:
	serv := createServerWithCreate(cr)
	retHasCreate := serv.HasCreate()
	retCreate := serv.GetCreate()
	retHasDelete := serv.HasDelete()

	if !retHasCreate {
		t.Errorf("the returned HasCreate has expected to be true, false returned")
	}

	if !reflect.DeepEqual(cr, retCreate) {
		t.Errorf("the returned create is invalid.")
	}

	if retHasDelete {
		t.Errorf("the returned HasDelete has expected to be false, true returned")
	}

}

func TestCreateServer_withDelete_Success(t *testing.T) {

	//variables:
	del := CreateDeleteForTests(t)

	//execute:
	serv := createServerWithDelete(del)
	retHasCreate := serv.HasCreate()
	retHasDelete := serv.HasDelete()
	retDelete := serv.GetDelete()

	if retHasCreate {
		t.Errorf("the returned HasCreate has expected to be false, true returned")
	}

	if !retHasDelete {
		t.Errorf("the returned HasDelete has expected to be true, false returned")
	}

	if !reflect.DeepEqual(del, retDelete) {
		t.Errorf("the returned create is invalid.")
	}

}

func TestCreateServer_withCreate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Server)
	obj := CreateServerWithCreateForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}

func TestCreateServer_withDelete_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Server)
	obj := CreateServerWithDeleteForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
