package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"reflect"
	"testing"

	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreatePublicKey_Success(t *testing.T) {

	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, keyErr := rsa.GenerateKey(reader, bitSize)

	if keyErr != nil {
		t.Errorf("there was a problem while generating an rsa.PrivateKey")
	}

	//execute:
	pk := createPublicKey(&key.PublicKey)

	retKey := pk.GetKey()
	retStr, retStrErr := pk.String()
	if retStrErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retStrErr.Error())
	}

	if !reflect.DeepEqual(&key.PublicKey, retKey) {
		t.Errorf("the returned public key was invalid")
	}

	//re-create the public key using the encoded string:
	otherPk, otherPkErr := createPublicKeyFromEncodedString(retStr)
	if otherPkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", otherPkErr.Error())
	}

	if !reflect.DeepEqual(pk, otherPk) {
		t.Errorf("the returned public key was invalid")
	}
}

func TestCreatePublicKey_convertToJSON_backAndForth_Success(t *testing.T) {

	//variables:
	empty := new(PublicKey)
	pk := CreatePublicKeyForTests(t)

	//execute:
	convert.ConvertToJSON(t, pk, empty)
}
