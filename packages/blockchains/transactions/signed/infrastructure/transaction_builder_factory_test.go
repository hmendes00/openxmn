package infrastructure

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	concrete_metadata "github.com/XMNBlockchain/core/packages/blockchains/metadata/infrastructure"
)

func TestCreateTransactionBuilderFactory_Success(t *testing.T) {

	//variables:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaBuilderFactory := concrete_metadata.CreateMetaDataBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metaBuilderFactory)

	//execute:
	fac := CreateTransactionBuilderFactory(htBuilderFactory, metaBuilderFactory)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
