package chunks

import (
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/hashtrees"
)

func TestCreateBuilder_Success(t *testing.T) {

	//variables:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	chkSizeInBytes := 3
	extension := "tmp"
	build := createBuilder(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)

	//execute:
	fac := CreateBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
