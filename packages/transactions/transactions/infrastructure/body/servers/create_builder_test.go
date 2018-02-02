package servers

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	concrete_servers "github.com/XMNBlockchain/core/packages/servers/infrastructure"
	uuid "github.com/satori/go.uuid"
)

func TestBuildCreate_Success(t *testing.T) {

	// random generator
	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)

	//variables:
	id := uuid.NewV4()
	serv := concrete_servers.CreateServerForTests(t)
	isTrs := r.Intn(2) != 0
	isLead := r.Intn(2) != 0
	isDB := r.Intn(2) != 0

	//to make sure there is at least one of them to true:
	if !isLead && !isDB {
		isTrs = true
	}

	//execute:
	build := createCreateBuilder()
	build.Create().WithID(&id).WithServer(serv)

	if isTrs {
		build.IsTransaction()
	}

	if isLead {
		build.IsLeader()
	}

	if isDB {
		build.IsDatabase()
	}

	cr, crErr := build.Now()

	if crErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", crErr.Error())
	}

	if cr == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

	retID := cr.GetID()
	retServ := cr.GetServer()
	retIsTrs := cr.IsTransaction()
	retIsLead := cr.IsLeader()
	retIsDB := cr.IsDatabase()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(serv, retServ) {
		t.Errorf("the returned server is invalid")
	}

	if isTrs != retIsTrs {
		t.Errorf("the returned IsTransaction is invalid.  Expected: %t, Returned: %t", isTrs, retIsTrs)
	}

	if isLead != retIsLead {
		t.Errorf("the returned IsLeader is invalid.  Expected: %t, Returned: %t", isLead, retIsLead)
	}

	if isDB != retIsDB {
		t.Errorf("the returned IsDatabase is invalid.  Expected: %t, Returned: %t", isDB, retIsDB)
	}

}

func TestBuildCreate_isTransaction_isLeader_isDatabase_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	serv := concrete_servers.CreateServerForTests(t)

	//execute:
	build := createCreateBuilder()
	cr, crErr := build.Create().WithID(&id).WithServer(serv).IsTransaction().IsLeader().IsDatabase().Now()

	if crErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", crErr.Error())
	}

	if cr == nil {
		t.Errorf("the returned instance was expected to be valid, nil returned")
	}

	retID := cr.GetID()
	retServ := cr.GetServer()
	retIsTrs := cr.IsTransaction()
	retIsLead := cr.IsLeader()
	retIsDB := cr.IsDatabase()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(serv, retServ) {
		t.Errorf("the returned server is invalid")
	}

	if !retIsTrs {
		t.Errorf("the returned IsTransaction is invalid.  Expected: %t, Returned: %t", true, retIsTrs)
	}

	if !retIsLead {
		t.Errorf("the returned IsLeader is invalid.  Expected: %t, Returned: %t", true, retIsLead)
	}

	if !retIsDB {
		t.Errorf("the returned IsDatabase is invalid.  Expected: %t, Returned: %t", true, retIsDB)
	}

}

func TestBuildCreate_withoutType_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	serv := concrete_servers.CreateServerForTests(t)

	//execute:
	build := createCreateBuilder()
	cr, crErr := build.Create().WithID(&id).WithServer(serv).Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildCreate_withoutID_returnsError(t *testing.T) {

	//variables:
	serv := concrete_servers.CreateServerForTests(t)

	//execute:
	build := createCreateBuilder()
	cr, crErr := build.Create().WithServer(serv).IsLeader().Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildCreate_withoutServer_returnsError(t *testing.T) {

	//variables:
	id := uuid.NewV4()

	//execute:
	build := createCreateBuilder()
	cr, crErr := build.Create().WithID(&id).IsLeader().Now()

	if crErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if cr != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
