package users

import (
	"reflect"
	"testing"
)

func TestBuildUser_withCreate_Success(t *testing.T) {

	//variables:
	cr := CreateSaveForTests(t)

	//execute:
	usr, usrErr := createUserBuilder().Create().WithSave(cr).Now()
	if usrErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", usrErr.Error())
	}

	if usr == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildUser_withDelete_Success(t *testing.T) {

	//variables:
	del := CreateDeleteForTests(t)

	//execute:
	usr, usrErr := createUserBuilder().Create().WithDelete(del).Now()
	if usrErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", usrErr.Error())
	}

	if usr == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildUser_withoutTrs_returnsError(t *testing.T) {

	//execute:
	usr, usrErr := createUserBuilder().Create().Now()
	if usrErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if usr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
