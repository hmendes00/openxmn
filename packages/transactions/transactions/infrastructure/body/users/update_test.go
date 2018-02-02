package users

import (
	"reflect"
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateUpdate_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests(t)

	//execute:
	up := createUpdate(&id, pk)

	retID := up.GetID()
	retPK := up.GetNewPublicKey()

	if !reflect.DeepEqual(&id, retID) {
		t.Errorf("the returned ID was invalid.  Expected: %s, Returned: %s", id.String(), retID.String())
	}

	if !reflect.DeepEqual(pk, retPK) {
		t.Errorf("the returned public key is invalid")
	}

}

func TestCreateUpdate_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Update)
	obj := CreateUpdateForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
