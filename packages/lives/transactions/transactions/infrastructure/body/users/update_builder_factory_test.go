package users

import (
	"reflect"
	"testing"
)

func TestUpdateBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createUpdateBuilder()

	//execute:
	fac := CreateUpdateBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
