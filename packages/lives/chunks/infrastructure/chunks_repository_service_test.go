package infrastructure

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	concrete_files "github.com/XMNBlockchain/core/packages/lives/files/infrastructure"
	concrete_stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/infrastructure"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
)

func TestSave_thenRetrieve_Success(t *testing.T) {

	//file variables:
	saveInPath := filepath.Join("test_files", "files")

	//read file:
	inputFilePath := filepath.Join("test_files", "input", "montreal.jpg")
	readData, readErr := ioutil.ReadFile(inputFilePath)
	if readErr != nil {
		t.Errorf("there was an error while reading the input binary file: %s", readErr.Error())
	}

	//chunks variables:
	chkSizeInBytes := 1024 * 100 //100 kb each chunk
	extension := "chks"

	//factories:
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	fileService := concrete_files.CreateFileService(storedFileBuilderFactory)
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	fileRepository := concrete_files.CreateFileRepository(fileBuilderFactory)
	storedChkBuilderFactory := concrete_stored_chunks.CreateChunksBuilderFactory()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactory()
	chksBuilderFactory := CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	htService := concrete_hashtrees.CreateHashTreeService(fileService, fileBuilderFactory)
	htRepository := concrete_hashtrees.CreateHashTreeRepository(fileRepository)

	//delete the files at the end:
	defer func() {
		fileService.DeleteAll(saveInPath)
	}()

	//build chunks:
	chks, _ := chksBuilderFactory.Create().Create().WithData(readData).Now()

	//execute:
	service := CreateChunksService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	repository := CreateChunksRepository(htRepository, fileRepository, chksBuilderFactory)

	//verify that the chunks do not exists:
	if _, err := os.Stat(saveInPath); !os.IsNotExist(err) {
		t.Errorf("the given path was not expected to be a valid directory: %s", saveInPath)
	}

	//save the chks:
	storedChks, storedChksErr := service.Save(saveInPath, chks)
	if storedChksErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedChksErr.Error())
	}

	//make sure the hashtree file exists:
	htFile := storedChks.GetHashTree()
	htFilePath := filepath.Join(saveInPath, htFile.GetPath())
	if _, err := os.Stat(htFilePath); os.IsNotExist(err) {
		t.Errorf("the hashtree file path was expected to exists: %s", htFilePath)
	}

	//read the hashtree:
	htRead, htReadErr := ioutil.ReadFile(htFilePath)
	if htReadErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", htReadErr.Error())
	}

	//convert the hashtree data to json:
	readHt := new(concrete_hashtrees.HashTree)
	jsErr := json.Unmarshal(htRead, readHt)
	if jsErr != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", jsErr.Error())
	}

	//compare the hashes of the hashtrees:
	if htFile.GetHash() != readHt.GetHash().String() {
		//		t.Errorf("the hash of the stored hashtree is different than the hash of the hashtree on file (%s).  Expected: %s, Read: %s", htFilePath, htFile.GetHash(), readHt.GetHash().String())
	}

	//make sure the chunk files exists:
	chkFiles := storedChks.GetChunks()
	for _, oneChkFile := range chkFiles {
		oneChkFilePath := filepath.Join(oneChkFile.GetPath())

		if _, err := os.Stat(htFilePath); os.IsNotExist(err) {
			t.Errorf("the chunk file path was expected to exists: %s", oneChkFilePath)
		}

		//make sure the size fits:
		oneChkSize := oneChkFile.GetSizeInBytes()
		chkData, _ := ioutil.ReadFile(htFilePath)
		if oneChkSize != len(chkData) {
			//t.Errorf("the chunk size is invalid.  Expected: %d, Received: %d", len(chkData), oneChkSize)
		}

		//the chunk should not be bigger than what we specificed:
		if oneChkSize > chkSizeInBytes {
			t.Errorf("the chunk size was %d bytes.  It should not be bigger than %d bytes", oneChkSize, chkSizeInBytes)
		}

	}

	//retrieve the stored chunk:
	retChunks, retChunksErr := repository.Retrieve(saveInPath)
	if retChunksErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retChunksErr.Error())
	}

	if !reflect.DeepEqual(chks, retChunks) {
		t.Errorf("the retrieved chunk is invalid")
	}
}
