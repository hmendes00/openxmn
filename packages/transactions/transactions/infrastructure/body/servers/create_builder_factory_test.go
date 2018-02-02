package servers

import (
	"reflect"
	"testing"
)

func TestCreateBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createCreateBuilder()

	//execute:
	fac := CreateCreateBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
