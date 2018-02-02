package infrastructure

import (
	"reflect"
	"testing"
)

func TestCreateServerBuilder_withURI_Success(t *testing.T) {

	//variables:
	build := createServerBuilder()

	//execute:
	fac := CreateServerBuilderFactory()
	retCreate := fac.Create()

	if !reflect.DeepEqual(build, retCreate) {
		t.Errorf("the returned builder is invalid")
	}

}
