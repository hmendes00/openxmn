package rsa

import (
	"reflect"
	"testing"
)

func TestCreatePrivateKeyBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createPrivateKeyBuilder()

	//execute:
	fac := CreatePrivateKeyBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder was invalid")
	}

}
