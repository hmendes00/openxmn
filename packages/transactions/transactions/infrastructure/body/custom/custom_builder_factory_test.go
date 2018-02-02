package custom

import (
	"reflect"
	"testing"
)

func TestCreateCustomBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createCustomBuilder()

	//execute:
	fac := CreateCustomBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
