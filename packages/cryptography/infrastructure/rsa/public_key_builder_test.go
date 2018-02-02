package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"reflect"
	"testing"
)

func TestBuildPublicKey_withKey_Success(t *testing.T) {

	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, keyErr := rsa.GenerateKey(reader, bitSize)

	if keyErr != nil {
		t.Errorf("there was a problem while generating an rsa.PrivateKey")
	}

	//execute:
	build := createPublicKeyBuilder()
	pk, pkErr := build.Create().WithKey(&key.PublicKey).Now()

	if pkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned")
	}

	retKey := pk.GetKey()

	if !reflect.DeepEqual(&key.PublicKey, retKey) {
		t.Errorf("the returned public key was invalid")
	}
}

func TestBuildPublicKey_withEncodedString_Success(t *testing.T) {

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
	build := createPublicKeyBuilder()
	otherPk, otherPkErr := build.Create().WithEncodedString(retStr).Now()
	if otherPkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", otherPkErr.Error())
	}

	if !reflect.DeepEqual(pk, otherPk) {
		t.Errorf("the returned public key was invalid")
	}
}

func TestBuildPublicKey_withKey_withEncodedString_returnsError(t *testing.T) {

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
	build := createPublicKeyBuilder()
	otherPk, otherPkErr := build.Create().WithKey(retKey).WithEncodedString(retStr).Now()
	if otherPkErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if otherPk != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}

func TestBuildPublicKey_withoutKey_withoutEncodedString_returnsError(t *testing.T) {

	//re-create the public key using the encoded string:
	build := createPublicKeyBuilder()
	pk, pkErr := build.Create().Now()
	if pkErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if pk != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}
}
