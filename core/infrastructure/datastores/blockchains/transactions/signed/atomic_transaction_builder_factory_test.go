package signed

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/metadata"
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
