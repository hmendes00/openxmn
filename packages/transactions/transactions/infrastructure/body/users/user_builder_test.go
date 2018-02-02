package users

import (
	"reflect"
	"testing"
)

func TestBuildUser_withCreate_Success(t *testing.T) {

	//variables:
	cr := CreateCreateForTests(t)

	//execute:
	usr, usrErr := createUserBuilder().Create().WithCreate(cr).Now()
	if usrErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", usrErr.Error())
	}

	if usr == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildUser_withUpdate_Success(t *testing.T) {

	//variables:
	up := CreateUpdateForTests(t)

	//execute:
	usr, usrErr := createUserBuilder().Create().WithUpdate(up).Now()
	if usrErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", usrErr.Error())
	}

	if usr == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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
