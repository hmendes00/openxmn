package infrastructure

import (
	"reflect"
	"testing"
)

func TestCreateTransactionsBuilderFactory_Success(t *testing.T) {

	//variables:
	build := createTransactionsBuilder()

	//execute:
	fac := CreateTransactionsBuilderFactory()
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
