package body

import (
	"reflect"
	"testing"

	concrete_custom "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body/custom"
	concrete_servers "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body/servers"
	concrete_users "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body/users"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateBody_withCustom_Success(t *testing.T) {

	//variables:
	cus := concrete_custom.CreateCustomWithCreateForTests(t)

	//execute:
	bod := createBodyWithCustom(cus)
	retHasCustom := bod.HasCustom()
	retCustom := bod.GetCustom()
	retHasServer := bod.HasServer()
	retHasUser := bod.HasUser()

	if !retHasCustom {
		t.Errorf("the returned HasCreate has expected to be true, false returned")
	}

	if !reflect.DeepEqual(cus, retCustom) {
		t.Errorf("the returned custom is invalid.")
	}

	if retHasServer {
		t.Errorf("the returned HasServer has expected to be false, true returned")
	}

	if retHasUser {
		t.Errorf("the returned HasUser has expected to be false, true returned")
	}

}

func TestCreateBody_withServer_Success(t *testing.T) {

	//variables:
	serv := concrete_servers.CreateServerWithCreateForTests(t)

	//execute:
	bod := createBodyWithServer(serv)
	retHasCustom := bod.HasCustom()
	retHasServer := bod.HasServer()
	retServer := bod.GetServer()
	retHasUser := bod.HasUser()

	if retHasCustom {
		t.Errorf("the returned HasCustom has expected to be false, true returned")
	}

	if !retHasServer {
		t.Errorf("the returned HasServer has expected to be true, false returned")
	}

	if !reflect.DeepEqual(serv, retServer) {
		t.Errorf("the returned server is invalid.")
	}

	if retHasUser {
		t.Errorf("the returned HasUser has expected to be false, true returned")
	}

}

func TestCreateBody_withUser_Success(t *testing.T) {

	//variables:
	usr := concrete_users.CreateUserWithSaveForTests(t)

	//execute:
	bod := createBodyWithUser(usr)
	retHasCustom := bod.HasCustom()
	retHasServer := bod.HasServer()
	retHasUser := bod.HasUser()
	retUser := bod.GetUser()

	if retHasCustom {
		t.Errorf("the returned HasCustom has expected to be false, true returned")
	}

	if retHasServer {
		t.Errorf("the returned HasServer has expected to be false, true returned")
	}

	if !retHasUser {
		t.Errorf("the returned HasUser has expected to be true, false returned")
	}

	if !reflect.DeepEqual(usr, retUser) {
		t.Errorf("the returned user is invalid.")
	}

}

func TestCreateUser_withCustom_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Body)
	obj := CreateBodyWithCustomForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}

func TestCreateUser_withServer_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Body)
	obj := CreateBodyWithServerForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}

func TestCreateUser_withUser_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Body)
	obj := CreateBodyWithUserForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
