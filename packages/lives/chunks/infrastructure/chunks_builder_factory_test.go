package infrastructure

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/lives/hashtrees/infrastructure"
	concrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
)

func TestCreateChunksBuilder_Success(t *testing.T) {

	//variables:
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	chkSizeInBytes := 3
	extension := "tmp"
	build := createChunksBuilder(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)

	//execute:
	fac := CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	retBuild := fac.Create()

	if !reflect.DeepEqual(build, retBuild) {
		t.Errorf("the returned TransactionBuilderFactory is invalid")
	}

}
