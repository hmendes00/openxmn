package servers

import (
	"reflect"
	"testing"
)

func TestCreateDeleteBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createDeleteBuilder()

	//execute:
	fac := CreateDeleteBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
