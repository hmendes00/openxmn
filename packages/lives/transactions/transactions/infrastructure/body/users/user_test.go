package users

import (
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateUser_withSave_Success(t *testing.T) {

	//variables:
	cr := CreateSaveForTests(t)

	//execute:
	usr := createUserWithSave(cr)
	retHasSave := usr.HasSave()
	retCreate := usr.GetSave()
	retHasDelete := usr.HasDelete()

	if !retHasSave {
		t.Errorf("the returned HasSave has expected to be true, false returned")
	}

	if !reflect.DeepEqual(cr, retCreate) {
		t.Errorf("the returned create is invalid.")
	}

	if retHasDelete {
		t.Errorf("the returned HasDelete has expected to be false, true returned")
	}

}

func TestCreateUser_withDelete_Success(t *testing.T) {

	//variables:
	del := CreateDeleteForTests(t)

	//execute:
	usr := createUserWithDelete(del)
	retHasSave := usr.HasSave()
	retHasDelete := usr.HasDelete()
	retDelete := usr.GetDelete()

	if retHasSave {
		t.Errorf("the returned HasSave has expected to be false, true returned")
	}

	if !retHasDelete {
		t.Errorf("the returned HasDelete has expected to be true, false returned")
	}

	if !reflect.DeepEqual(del, retDelete) {
		t.Errorf("the returned delete is invalid.")
	}

}

func TestCreateUser_withSave_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(User)
	obj := CreateUserWithSaveForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}

func TestCreateUser_withDelete_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(User)
	obj := CreateUserWithDeleteForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
