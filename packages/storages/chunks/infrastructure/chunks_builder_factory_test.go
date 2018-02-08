package infrastructure

import (
	"reflect"
	"testing"
)

func TestCreateChunksBuilder_Success(t *testing.T) {

	//variables:
	build := createChunksBuilder()

	//execute:
	fac := CreateChunksBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
