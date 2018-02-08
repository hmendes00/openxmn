package infrastructure

import (
	"reflect"
	"testing"
	"time"

	files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	concrete_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
)

func TestBuildChunks_Success(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests(t)
	chksFiles := []files.File{
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
	}
	createdOn := time.Now().UTC()

	//execute:
	build := createChunksBuilder()
	chks, chksErr := build.Create().WithChunks(chksFiles).WithHashTree(htFile).CreatedOn(createdOn).Now()
	if chksErr != nil {
		t.Errorf("the returned error was expected to be nil, returned: %s", chksErr.Error())
	}

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

func TestBuildChunks_withoutCreatedOn_returnsError(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests(t)
	chksFiles := []files.File{
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
	}

	//execute:
	build := createChunksBuilder()
	chks, chksErr := build.Create().WithChunks(chksFiles).WithHashTree(htFile).Now()
	if chksErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if chks != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildChunks_withoutHashTree_returnsError(t *testing.T) {

	//variables:
	chksFiles := []files.File{
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
		concrete_files.CreateFileForTests(t),
	}
	createdOn := time.Now().UTC()

	//execute:
	build := createChunksBuilder()
	chks, chksErr := build.Create().WithChunks(chksFiles).CreatedOn(createdOn).Now()
	if chksErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if chks != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildChunks_withoutChunks_returnsError(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests(t)
	createdOn := time.Now().UTC()

	//execute:
	build := createChunksBuilder()
	chks, chksErr := build.Create().WithHashTree(htFile).CreatedOn(createdOn).Now()
	if chksErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if chks != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}

func TestBuildChunks_withEmptyChunks_returnsError(t *testing.T) {

	//variables:
	htFile := concrete_files.CreateFileForTests(t)
	chksFiles := []files.File{}
	createdOn := time.Now().UTC()

	//execute:
	build := createChunksBuilder()
	chks, chksErr := build.Create().WithChunks(chksFiles).WithHashTree(htFile).CreatedOn(createdOn).Now()
	if chksErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	if chks != nil {
		t.Errorf("the returned instance was expected to be nil, instance returned")
	}

}
