package servers

import (
	"reflect"
	"testing"
)

func TestBuildServer_withCreate_Success(t *testing.T) {

	//variables:
	cr := CreateCreateForTests(t)

	//execute:
	serv, servErr := createServerBuilder().Create().WithCreate(cr).Now()
	if servErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", servErr.Error())
	}

	if serv == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildServer_withDelete_Success(t *testing.T) {

	//variables:
	del := CreateDeleteForTests(t)

	//execute:
	serv, servErr := createServerBuilder().Create().WithDelete(del).Now()
	if servErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", servErr.Error())
	}

	if serv == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildServer_withoutTrs_returnsError(t *testing.T) {

	//execute:
	serv, servErr := createServerBuilder().Create().Now()
	if servErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if serv != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
