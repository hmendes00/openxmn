package files

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRetrieve_Success(t *testing.T) {

	//variables
	dat := []byte("this is some data.  Lets see how it goes.")
	sizeInBytes := len(dat)
	filePath := "test_files/myfile.tmp"
	rep := CreateFileRepositoryForTests()
	defer os.Remove(filePath)

	//try to retrieve the file, should not work:
	_, doNotWorkFilErr := rep.Retrieve(filePath)
	if doNotWorkFilErr == nil {
		t.Errorf("the returned error was expected to be valid, nil returned")
	}

	//save the file:
	err := ioutil.WriteFile(filePath, dat, 0777)
	if err != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", err.Error())
	}

	//retrieve the file, should work:
	fil, filErr := rep.Retrieve(filePath)
	if filErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", filErr.Error())
	}

	opFile, _ := os.Open(filePath)
	statFile, _ := opFile.Stat()

	retPath := fil.GetPath()
	retCreatedOn := fil.CreatedOn()
	retSizeInBytes := fil.GetSizeInBytes()

	if filePath != retPath {
		t.Errorf("the returned path is invalid.  Expected: %s, Returned: %s", filePath, retPath)
	}

	if statFile.ModTime() != retCreatedOn {
		t.Errorf("the returned createdOn is invalid.  Expected: %s, Returned: %s", statFile.ModTime(), retCreatedOn)
	}

	if sizeInBytes != retSizeInBytes {
		t.Errorf("the returned sizeInBytes is invalid.  Expected: %d, Returned: %d", sizeInBytes, retSizeInBytes)
	}
}

func TestRetrieveAll_Success(t *testing.T) {

	//variables:
	dirPath := "test_files"
	firstDat := []byte("this is some data.  Lets see how it goes.")
	firstFilePath := filepath.Join(dirPath, "my_first_file.tmp")
	secondDat := []byte("some other data")
	secondFilePath := filepath.Join(dirPath, "my_second_file.tmp")
	rep := CreateFileRepositoryForTests()
	defer func() {
		os.Remove(firstFilePath)
		os.Remove(secondFilePath)
	}()

	//try to retrieve the file, should be empty:
	emptyFiles, emptyFilesErr := rep.RetrieveAll(dirPath)
	if emptyFilesErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", emptyFilesErr.Error())
	}

	if len(emptyFiles) != 0 {
		t.Errorf("there should be no files in the directory")
	}

	//save the files:
	ioutil.WriteFile(firstFilePath, firstDat, 0777)
	ioutil.WriteFile(secondFilePath, secondDat, 0777)

	//retrieve the files, should work:
	fil, filErr := rep.RetrieveAll(dirPath)
	if filErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", filErr.Error())
	}

	if len(fil) != 2 {
		t.Errorf("there should be 2 files in the directory.  Found: %d", len(fil))
	}
}
