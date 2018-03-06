package infrastructure

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
)

func TestCreateSignedTransactionsBuilderFactory_Success(t *testing.T) {

	//variables:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	build := createSignedTransactionsBuilder(htBuilderFactory, metaDataBuilderFactory)

	//execute:
	fac := CreateSignedTransactionsBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
