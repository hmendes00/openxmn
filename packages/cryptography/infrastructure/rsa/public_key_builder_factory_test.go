package rsa

import (
	"reflect"
	"testing"
)

func TestCreatePublicKeyBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createPublicKeyBuilder()

	//execute:
	fac := CreatePublicKeyBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder was invalid")
	}

}
