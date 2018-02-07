package infrastructure

import (
	"reflect"
	"testing"
)

func TestCreateFileBuilder_Success(t *testing.T) {

	//variables:
	build := createFileBuilder()

	//execute:
	fac := CreateFileBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
