package servers

import (
	"reflect"
	"testing"
)

func TestCreateServerBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createServerBuilder()

	//execute:
	fac := CreateServerBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
