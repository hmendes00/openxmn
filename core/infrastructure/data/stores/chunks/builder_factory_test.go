package chunks

import (
	"reflect"
	"testing"
)

func TestCreateChunksBuilder_Success(t *testing.T) {

	//variables:
	build := createBuilder()

	//execute:
	fac := CreateBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
