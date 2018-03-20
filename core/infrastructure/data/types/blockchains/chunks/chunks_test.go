package chunks

import (
	"reflect"
	"testing"

	files "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/files"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/files"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/hashtrees"
)

func TestCreateChunks_Success(t *testing.T) {

	//variables:
	chksFiles := []files.File{
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
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
