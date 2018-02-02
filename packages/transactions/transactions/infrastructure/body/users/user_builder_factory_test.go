package users

import (
	"reflect"
	"testing"
)

func TestCreateUserBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createUserBuilder()

	//execute:
	fac := CreateUserBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
