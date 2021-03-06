package chunks

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
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
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	chksBuilderFactory := CreateBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)

	//delete the files at the end:
	defer func() {
		fileService.DeleteAll(saveInPath)
	}()

	//build chunks:
	chks, _ := chksBuilderFactory.Create().Create().WithData(readData).Now()

	//execute:
	service := CreateServiceForTests()
	repository := CreateRepositoryForTests()

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
