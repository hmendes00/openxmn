package infrastructure

import (
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/infrastructure"
	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
	concrete_files "github.com/XMNBlockchain/core/packages/blockchains/files/infrastructure"
)

func TestCreateChunks_Success(t *testing.T) {

	//variables:
	chksFiles := []files.File{
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
	}

	blocks := [][]byte{}
	for _, oneChkFile := range chksFiles {
		blocks = append(blocks, oneChkFile.GetHash().Sum(nil))
	}

	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()

	//execute:
	chks := createChunks(ht, chksFiles)
	retHt := chks.GetHashTree()
	retChksFiles := chks.GetChunks()

	if !reflect.DeepEqual(chksFiles, retChksFiles) {
		t.Errorf("the returned chunks files are invalid")
	}

	if !reflect.DeepEqual(ht, retHt) {
		t.Errorf("the returned hashtree is invalid")
	}

}
