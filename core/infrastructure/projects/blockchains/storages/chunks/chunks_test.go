package chunks

import (
	"reflect"
	"testing"
	"time"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
	convert "github.com/XMNBlockchain/exmachina-network/core/infrastructure/tests/jsonify/helpers"
)

func TestCreateChunks_Success(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests()
	chksFiles := []*concrete_files.File{
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
	}
	createdOn := time.Now().UTC()

	//execute:
	chks := createChunks(htFile, chksFiles, createdOn)
	retHt := chks.GetHashTree()
	retChks := chks.GetChunks()
	retCreatedOn := chks.CreatedOn()

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

	if createdOn != retCreatedOn {
		t.Errorf("the returned creation time is invalid")
	}

}

func TestCreateFile_convertToJS_convertToInstance_Success(t *testing.T) {

	//variables:
	empty := new(Chunks)
	obj := CreateChunksForTests()

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
