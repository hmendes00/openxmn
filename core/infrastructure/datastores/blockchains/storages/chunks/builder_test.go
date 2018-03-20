package chunks

import (
	"reflect"
	"testing"

	files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
)

func TestBuildChunks_Success(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests()
	chksFiles := []files.File{
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
	}

	//execute:
	build := createBuilder()
	chks, chksErr := build.Create().WithChunks(chksFiles).WithHashTree(htFile).Now()
	if chksErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", chksErr.Error())
	}

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

func TestBuildChunks_withoutHashTree_returnsError(t *testing.T) {

	//variables:
	chksFiles := []files.File{
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
		concrete_files.CreateFileForTests(),
	}

	//execute:
	build := createBuilder()
	chks, chksErr := build.Create().WithChunks(chksFiles).Now()
	if chksErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if chks != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildChunks_withoutChunks_returnsError(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests()

	//execute:
	build := createBuilder()
	chks, chksErr := build.Create().WithHashTree(htFile).Now()
	if chksErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if chks != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildChunks_withEmptyChunks_returnsError(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests()
	chksFiles := []files.File{}

	//execute:
	build := createBuilder()
	chks, chksErr := build.Create().WithChunks(chksFiles).WithHashTree(htFile).Now()
	if chksErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if chks != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
