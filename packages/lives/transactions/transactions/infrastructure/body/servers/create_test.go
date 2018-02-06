package servers

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	concrete_servers "github.com/XMNBlockchain/core/packages/servers/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateCreate_Success(t *testing.T) {

	// random generator
	var src = rand.NewSource(time.Now().UnixNano())
	var r = rand.New(src)

	//variables:
	id := uuid.NewV4()
	serv := concrete_servers.CreateServerForTests(t)
	isTrs := r.Intn(2) != 0
	isLead := r.Intn(2) != 0
	isDB := r.Intn(2) != 0

	//execute:
	cr := createCreate(&id, serv, isTrs, isLead, isDB)

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

func TestCreateCreate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Create)
	obj := CreateCreateForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
