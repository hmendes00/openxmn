package infrastructure

import (
	"reflect"
	"testing"

	concrete_cryptography "github.com/XMNBlockchain/core/packages/cryptography/infrastructure/rsa"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateSignature_Success(t *testing.T) {

	//variables:
	sig := concrete_cryptography.CreateSignatureForTests(t)
	usr := CreateUserForTests(t)

	//execute:
	userSig := createSignature(sig, usr)

	retSig := userSig.GetSig()
	retUser := userSig.GetUser()

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
	obj := CreateSignatureForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
