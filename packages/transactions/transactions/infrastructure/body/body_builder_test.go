package body

import (
	"reflect"
	"testing"

	concrete_custom "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/custom"
	concrete_servers "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/servers"
	concrete_users "github.com/XMNBlockchain/core/packages/transactions/transactions/infrastructure/body/users"
)

func TestBuildBody_withCustom_Success(t *testing.T) {

	//variables:
	cus := concrete_custom.CreateCustomWithCreateForTests(t)

	//execute:
	bod, bodErr := createBodyBuilder().Create().WithCustom(cus).Now()
	if bodErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", bodErr.Error())
	}

	if bod == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildBody_withServer_Success(t *testing.T) {

	//variables:
	serv := concrete_servers.CreateServerWithCreateForTests(t)

	//execute:
	bod, bodErr := createBodyBuilder().Create().WithServer(serv).Now()
	if bodErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", bodErr.Error())
	}

	if bod == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildBody_withUser_Success(t *testing.T) {

	//variables:
	usr := concrete_users.CreateUserWithCreateForTests(t)

	//execute:
	bod, bodErr := createBodyBuilder().Create().WithUser(usr).Now()
	if bodErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", bodErr.Error())
	}

	if bod == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

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

func TestBuildBody_withoutTrs_Success(t *testing.T) {

	//execute:
	bod, bodErr := createBodyBuilder().Create().Now()
	if bodErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if bod != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
