package signed

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_metadata "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/blockchains/metadata"
)

func TestCreateTransactionBuilderFactory_Success(t *testing.T) {

	//variables:
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	metaBuilderFactory := concrete_metadata.CreateBuilderFactory()
	build := createTransactionBuilder(htBuilderFactory, metaBuilderFactory)

	//execute:
	fac := CreateTransactionBuilderFactory(htBuilderFactory, metaBuilderFactory)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned builder is invalid")
	}

}
