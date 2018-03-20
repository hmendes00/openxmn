package rsa

import (
	"reflect"
	"testing"
)

func TestCreateCipherBuilderFactory_Success(t *testing.T) {

	//variables:
	pk := CreatePrivateKeyForTests()
	publicKeyBuilderFactory := createPublicKeyBuilder()
	sigBuilderFactory := createSignatureBuilder(publicKeyBuilderFactory)
	build := createCipherBuilder(sigBuilderFactory, pk)

	//execute:
	fac := CreateCipherBuilderFactory(sigBuilderFactory, pk)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder was invalid")
	}

}
