package users

import (
	"reflect"
	"testing"
)

func TestCreateBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createSaveBuilder()

	//execute:
	fac := CreateSaveBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
