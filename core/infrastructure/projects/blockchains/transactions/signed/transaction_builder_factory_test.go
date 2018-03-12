package signed

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/metadata"
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
