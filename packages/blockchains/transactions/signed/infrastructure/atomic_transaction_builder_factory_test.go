package infrastructure

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
)

func TestCreateBuilderFactory_Success(t *testing.T) {

	//factories:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaDataBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()

	//variables:
	build := createAtomicTransactionBuilder(htBuilderFactory, metaDataBuilderFactory)

	//execute:
	fac := CreateAtomicTransactionBuilderFactory(htBuilderFactory, metaDataBuilderFactory)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
