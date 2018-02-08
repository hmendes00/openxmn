package infrastructure

import (
	"reflect"
	"testing"
	"time"

	concrete_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
	convert "github.com/XMNBlockchain/core/packages/tests/jsonify/helpers"
)

func TestCreateChunks_Success(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests(t)
	chksFiles := []*concrete_files.File{
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
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
	obj := CreateChunksForTests(t)

	//execute:
	convert.ConvertToJSON(t, obj, empty)
}
