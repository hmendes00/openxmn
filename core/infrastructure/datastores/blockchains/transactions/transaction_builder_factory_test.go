package transactions

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/hashtrees"
	concrete_met "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
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
