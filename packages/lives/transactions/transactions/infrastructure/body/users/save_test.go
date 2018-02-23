package users

import (
	"reflect"
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestSaveSave_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)

	//execute:
	cr := createSave(&id, pk)

	retID := cr.GetID()
	retPK := cr.GetPublicKey()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(pk, retPK) {
		t.Errorf("the returned public key is invalid")
	}

}

func TestSaveSave_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Save)
	obj := CreateSaveForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
