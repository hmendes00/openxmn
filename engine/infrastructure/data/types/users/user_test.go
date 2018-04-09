package users

import (
	"reflect"
	"testing"
	"time"

	concrete_cryptography "github.com/XMNBlockchain/openxmn/engine/infrastructure/cryptography"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/metadata"
	convert "github.com/XMNBlockchain/openxmn/engine/infrastructure/tests/jsonify/helpers"
	uuid "github.com/satori/go.uuid"
)

func TestCreateUser_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	pk := concrete_cryptography.CreatePublicKeyForTests()
	crOn := time.Now().UTC()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).CreatedOn(crOn).Now()

	//execute:
	user := createUser(met.(*concrete_metadata.MetaData), pk)

	retMetaData := user.GetMetaData()
	retPK := user.GetPublicKey()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned MetaData is invalid")
	}

	if !reflect.DeepEqual(pk, retPK) {
		t.Errorf("the returned public key is invalid")
	}
}

func TestCreateUser_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(User)
	obj := CreateUserForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
