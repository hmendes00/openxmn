package infrastructure

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_met "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
)

func TestCreateBuilderFactory_Success(t *testing.T) {

	//variables:
	metBuilderFactory := concrete_met.CreateMetaDataBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metBuilderFactory)

	//execute:
	fac := CreateTransactionBuilderFactory(htBuilderFactory, metBuilderFactory)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
