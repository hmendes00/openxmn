package chunks

import (
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/data/stores/files"
	convert "github.com/XMNBlockchain/exmachina-network/engine/infrastructure/tests/jsonify/helpers"
)

func TestCreateChunks_Success(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests()
	chksFiles := []*concrete_files.File{
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
	}

	//execute:
	chks := createChunks(htFile, chksFiles)
	retHt := chks.GetHashTree()
	retChks := chks.GetChunks()

	if !reflect.DeepEqual(htFile, retHt.(*concrete_files.File)) {
		t.Errorf("the returned hashtree file was invalid")
	}

	if len(chksFiles) != len(retChks) {
		t.Errorf("the amount of chunk files is invalid.  Expected: %d, Returned: %d", len(chksFiles), len(retChks))
	}

	for index, oneChks := range chksFiles {
		if !reflect.DeepEqual(oneChks, retChks[index].(*concrete_files.File)) {
			t.Errorf("the chunks file (index: %d) is invalid", index)
		}
	}

}

func TestCreateFile_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Chunks)
	obj := CreateChunksForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
