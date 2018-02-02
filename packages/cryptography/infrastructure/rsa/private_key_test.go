package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"reflect"
	"testing"
)

func TestCreatePrivateKey_Success(t *testing.T) {

	//variables:
	bitSize := 4096
	reader := rand.Reader
	key, keyErr := rsa.GenerateKey(reader, bitSize)

	if keyErr != nil {
		t.Errorf("there was a problem while generating an rsa.PrivateKey")
	}

	//execute:
	pk := createPrivateKey(key)

	retKey := pk.GetKey()

	if !reflect.DeepEqual(key, retKey) {
		t.Errorf("the returned private key was invalid")
	}

}
