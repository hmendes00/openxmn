package infrastructure

import (
	"reflect"
	"testing"
)

func TestCreateBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createAtomicTransactionBuilder()

	//execute:
	fac := CreateAtomicTransactionBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
