package aggregated

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/metadata"
)

func TestCreateTransactionsBuilderFactory_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	//variables:
	build := createTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)

	//execute:
	fac := CreateTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
