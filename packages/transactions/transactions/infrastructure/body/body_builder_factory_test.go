package body

import (
	"reflect"
	"testing"
)

func TestCreateBodyBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createBodyBuilder()

	//execute:
	fac := CreateBodyBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
