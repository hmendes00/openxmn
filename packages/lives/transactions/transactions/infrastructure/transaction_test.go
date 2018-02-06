package infrastructure

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	concrete_body "github.com/XMNBlockchain/core/packages/lives/transactions/transactions/infrastructure/body"
	uuid "github.com/satori/go.uuid"
)

func TestCreate_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	createdOn := time.Now()
	karma := rand.Int() % 20
	bod := concrete_body.CreateBodyWithCustomForTests(t)

	//execute:
	trs := createTransaction(&id, karma, bod, createdOn)
	retID := trs.GetID()
	retBody := trs.GetBody()
	retCreatedOn := trs.CreatedOn()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned id was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(bod, retBody) {
		t.Errorf("the returned body was invalid")
	}

	if !reflect.DeepEqual(createdOn, retCreatedOn) {
		t.Errorf("the returned createdOn was invalid")
	}
}

func TestCreateTransaction_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Transaction)
	obj := CreateTransactionForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
