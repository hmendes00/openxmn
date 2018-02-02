package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"reflect"
	"testing"
)

func TestBuildPrivateKey_withKey_Success(t *testing.T) {

	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, keyErr := rsa.GenerateKey(reader, bitSize)

	if keyErr != nil {
		t.Errorf("there was a problem while generating an rsa.PrivateKey")
	}

	//execute:
	build := createPrivateKeyBuilder()
	pk, pkErr := build.Create().WithKey(key).Now()

	if pkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", pkErr.Error())
	}

	retKey := pk.GetKey()

	if !reflect.DeepEqual(key, retKey) {
		t.Errorf("the returned private key was invalid")
	}

}

func TestBuildPrivateKey_withEncodedString_Success(t *testing.T) {

	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, keyErr := rsa.GenerateKey(reader, bitSize)
	if keyErr != nil {
		t.Errorf("there was a problem while generating an rsa.PrivateKey")
	}

	pk := createPrivateKey(key)
	pkAsString := pk.String()

	//execute:
	build := createPrivateKeyBuilder()
	pk, pkErr := build.Create().WithEncodedString(pkAsString).Now()

	if pkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", pkErr.Error())
	}

	retKey := pk.GetKey()

	if !reflect.DeepEqual(key, retKey) {
		t.Errorf("the returned private key was invalid")
	}

}

func TestBuildPrivateKey_withoutKey_withoutEncodedString_Success(t *testing.T) {
	//execute:
	build := createPrivateKeyBuilder()
	pk, pkErr := build.Create().Now()

	if pkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", pkErr.Error())
	}

	retKey := pk.GetKey()
	retStr := pk.String()

	//build with encoded string, using the string:
	otherPk, otherPkErr := build.Create().WithEncodedString(retStr).Now()

	if otherPkErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", otherPkErr.Error())
	}

	otherRetKey := otherPk.GetKey()

	if !reflect.DeepEqual(otherRetKey, retKey) {
		t.Errorf("the returned private key was invalid")
	}

}

func TestBuildPrivateKey_withKey_withEncodedString_Success(t *testing.T) {

	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, keyErr := rsa.GenerateKey(reader, bitSize)

	if keyErr != nil {
		t.Errorf("there was a problem while generating an rsa.PrivateKey")
	}

	pk := createPrivateKey(key)
	pkAsString := pk.String()

	//execute:
	build := createPrivateKeyBuilder()
	builtPk, builtPkErr := build.Create().WithKey(key).WithEncodedString(pkAsString).Now()

	if builtPkErr == nil {
		t.Errorf("the returned error was expected to be an error, nil returned")
	}

	if builtPk != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
