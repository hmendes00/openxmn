package infrastructure

import (
	"reflect"
	"testing"

	concrete_users "github.com/XMNBlockchain/core/packages/blockchains/users/infrastructure"
)

func TestCreateSignedTransactionsBuilderFactory_Success(t *testing.T) {

	//variables:
	usrSigBuilderFactory := concrete_users.CreateSignatureBuilderFactoryForTests()
	build := createSignedTransactionsBuilder(usrSigBuilderFactory)

	//execute:
	fac := CreateSignedTransactionsBuilderFactory(usrSigBuilderFactory)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
