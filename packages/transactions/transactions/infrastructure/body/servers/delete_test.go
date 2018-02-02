package servers

import (
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateDelete_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()

	//execute:
	del := createDelete(&id)

	retID := del.GetID()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

}

func TestCreateDelete_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Delete)
	obj := CreateDeleteForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
