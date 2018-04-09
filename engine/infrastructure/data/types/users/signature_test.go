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

func TestCreateSignature_Success(t *testing.T) {

	//variables:
	id := uuid.NewV4()
	sig := concrete_cryptography.CreateSignatureForTests()
	usr := CreateUserUsingProvidedPublicKeyForTests(sig.GetPublicKey())
	crOn := time.Now().UTC()
	met, _ := concrete_metadata.CreateBuilderFactory().Create().Create().WithID(&id).CreatedOn(crOn).Now()

	//execute:
	userSig, userSigErr := createSignature(met.(*concrete_metadata.MetaData), sig, usr)
	if userSigErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", userSigErr.Error())
	}

	retMetaData := userSig.GetMetaData()
	retSig := userSig.GetSignature()
	retUser := userSig.GetUser()

	if !reflect.DeepEqual(met, retMetaData) {
		t.Errorf("the returned metadata was invalid")
	}

	if !reflect.DeepEqual(sig, retSig) {
		t.Errorf("the returned signature was invalid")
	}

	if !reflect.DeepEqual(usr, retUser) {
		t.Errorf("the user was invalid")
	}
}

func TestCreateSignature_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Signature)
	obj := CreateSignatureForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
