package infrastructure

import (
	"reflect"
	"testing"
)

func TestCreateBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createTransactionBuilder()

	//execute:
	fac := CreateTransactionBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
