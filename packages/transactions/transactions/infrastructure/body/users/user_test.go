package users

import (
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateUser_withCreate_Success(t *testing.T) {

	//variables:
	cr := CreateCreateForTests(t)

	//execute:
	usr := createUserWithCreate(cr)
	retHasCreate := usr.HasCreate()
	retCreate := usr.GetCreate()
	retHasDelete := usr.HasDelete()

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

func TestCreateUser_withDelete_Success(t *testing.T) {

	//variables:
	del := CreateDeleteForTests(t)

	//execute:
	usr := createUserWithDelete(del)
	retHasCreate := usr.HasCreate()
	retHasDelete := usr.HasDelete()
	retDelete := usr.GetDelete()

	if retHasCreate {
		t.Errorf("the returned HasCreate has expected to be false, true returned")
	}

	if !retHasDelete {
		t.Errorf("the returned HasDelete has expected to be true, false returned")
	}

	if !reflect.DeepEqual(del, retDelete) {
		t.Errorf("the returned delete is invalid.")
	}

}

func TestCreateUser_withUpdate_Success(t *testing.T) {

	//variables:
	up := CreateUpdateForTests(t)

	//execute:
	usr := createUserWithUpdate(up)
	retHasCreate := usr.HasCreate()
	retHasDelete := usr.HasDelete()
	retHasUpdate := usr.HasUpdate()
	retUpdate := usr.GetUpdate()

	if retHasCreate {
		t.Errorf("the returned HasCreate has expected to be false, true returned")
	}

	if retHasDelete {
		t.Errorf("the returned HasDelete has expected to be false, true returned")
	}

	if !retHasUpdate {
		t.Errorf("the returned HasUpdate has expected to be true, false returned")
	}

	if !reflect.DeepEqual(up, retUpdate) {
		t.Errorf("the returned update is invalid.")
	}

}

func TestCreateUser_withCreate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(User)
	obj := CreateUserWithCreateForTests(t)

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

func TestCreateUser_withUpdate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(User)
	obj := CreateUserWithUpdateForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
