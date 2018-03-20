package hashtrees

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/blockchains/files"
)

func TestSaveHashTree_thenRetrieve_Success(t *testing.T) {

	//file variables:
	saveInPath := filepath.Join("test_files", "files")

	//variables:
	htBlocks := [][]byte{
		[]byte("some"),
		[]byte("data"),
	}

	//factories:
	fileService := concrete_files.CreateFileServiceForTests()
	htBuilderFactory := CreateHashTreeBuilderFactoryForTests()

	//delete the files at the end:
	defer func() {
		fileService.DeleteAll(saveInPath)
	}()

	//build hashtree:
	ht, _ := htBuilderFactory.Create().Create().WithBlocks(htBlocks).Now()

	//execute:
	service := CreateHashTreeServiceForTests()
	repository := CreateHashTreeRepositoryForTests()

	//verify that the chunks do not exists:
	if _, err := os.Stat(saveInPath); !os.IsNotExist(err) {
		t.Errorf("the given path was not expected to be a valid directory: %s", saveInPath)
	}

	//save the hashtree on disk:
	storedHt, storedHtErr := service.Save(saveInPath, ht)
	if storedHtErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", storedHtErr.Error())
	}

	//read the hashtree:
	htFilePath := filepath.Join(saveInPath, storedHt.GetPath())
	htRead, htReadErr := ioutil.ReadFile(htFilePath)
	if htReadErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", htReadErr.Error())
	}

	//convert the hashtree data to json:
	readHt := new(HashTree)
	jsErr := json.Unmarshal(htRead, readHt)
	if jsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", jsErr.Error())
	}

	//compare the hashes of the hashtrees:
	if !reflect.DeepEqual(ht, readHt) {
		t.Errorf("the hash of the stored hashtree is different than the hash of the hashtree on file (%s).  Expected: %s, Read: %s", htFilePath, ht.GetHash(), readHt.GetHash().String())
	}

	//read the hashtree using the repository:
	retRepHt, retRepHtErr := repository.Retrieve(saveInPath)
	if retRepHtErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", retRepHtErr.Error())
	}

	//compare the hashes of the hashtrees:
	if !reflect.DeepEqual(ht, retRepHt) {
		t.Errorf("the hash of the stored hashtree is different than the hash of the hashtree on file (%s).  Expected: %s, Read: %s", htFilePath, ht.GetHash(), retRepHt.GetHash().String())
	}
}
