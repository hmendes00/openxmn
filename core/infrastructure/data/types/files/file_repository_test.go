package files

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	files "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/files"
)

func verifyFilesInRepositoryForTests(t *testing.T, expectedFiles []files.File, retrievedFiles []files.File) {
	for index, oneExpectedFile := range expectedFiles {
		retDirPath := retrievedFiles[index].GetDirPath()
		retFileName := retrievedFiles[index].GetFileName()
		retExt := retrievedFiles[index].GetExtension()
		retData := retrievedFiles[index].GetData()

		if retDirPath != oneExpectedFile.GetDirPath() {
			t.Errorf("the returned dirPath is invalid.  Expected: %s, Returned: %s", oneExpectedFile.GetDirPath(), retDirPath)
		}

		if retFileName != oneExpectedFile.GetFileName() {
			t.Errorf("the returned dirPath is invalid.  Expected: %s, Returned: %s", oneExpectedFile.GetFileName(), retFileName)
		}

		if retExt != oneExpectedFile.GetExtension() {
			t.Errorf("the returned extension is invalid.  Expected: %s, Returned: %s", oneExpectedFile.GetExtension(), retExt)
		}

		if !bytes.Equal(retData, oneExpectedFile.GetData()) {
			t.Errorf("the returned data is invalid")
		}
	}
}

func TestRetrieve_Success(t *testing.T) {

	//factories:
	fileBuilderFactory := CreateFileBuilderFactory()

	//file variabkles:
	retrieveInPath := filepath.Join("test_files", "input")
	fileName := "montreal"
	ext := "jpg"
	fileNameWithExt := fmt.Sprintf("%s.%s", fileName, ext)

	//read data:
	filePath := filepath.Join(retrieveInPath, fileNameWithExt)
	readData, readErr := ioutil.ReadFile(filePath)
	if readErr != nil {
		t.Errorf("there was an error while reading the input binary file: %s", readErr.Error())
	}

	//file:
	fil, _ := fileBuilderFactory.Create().Create().WithDirPath(retrieveInPath).WithFileName(fileName).WithExtension(ext).WithData(readData).Now()
	expectedFiles := []files.File{
		fil,
	}

	//execute:
	rep := CreateFileRepositoryForTests()
	retFile, retFileErr := rep.Retrieve(retrieveInPath, fileNameWithExt)
	if retFileErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retFileErr.Error())
	}

	retFiles := []files.File{
		retFile,
	}

	//verify:
	verifyFilesInRepositoryForTests(t, expectedFiles, retFiles)
}

func TestRetrieveAll_Success(t *testing.T) {

	//factories:
	fileBuilderFactory := CreateFileBuilderFactory()

	//file variabkles:
	retrieveInPath := filepath.Join("test_files", "input")
	ext := "jpg"

	//first read data:
	firstFileName := "montreal"
	firstFileNameWithExt := fmt.Sprintf("%s.%s", firstFileName, ext)
	firstFilePath := filepath.Join(retrieveInPath, firstFileNameWithExt)
	firstReadData, firstReadDataErr := ioutil.ReadFile(firstFilePath)
	if firstReadDataErr != nil {
		t.Errorf("there was an error while reading the input binary file: %s", firstReadDataErr.Error())
	}

	//second read data:
	secondFileName := "montreal_second"
	secondFileNameWithExt := fmt.Sprintf("%s.%s", secondFileName, ext)
	secondFilePath := filepath.Join(retrieveInPath, secondFileNameWithExt)
	secondReadData, secondReadDataErr := ioutil.ReadFile(secondFilePath)
	if secondReadDataErr != nil {
		t.Errorf("there was an error while reading the input binary file: %s", secondReadDataErr.Error())
	}

	//create the file names:
	fileNamesWithExt := []string{
		firstFileNameWithExt,
		secondFileNameWithExt,
	}

	//files:
	firstFile, _ := fileBuilderFactory.Create().Create().WithDirPath(retrieveInPath).WithFileName(firstFileName).WithExtension(ext).WithData(firstReadData).Now()
	secondFile, _ := fileBuilderFactory.Create().Create().WithDirPath(retrieveInPath).WithFileName(secondFileName).WithExtension(ext).WithData(secondReadData).Now()
	expectedFiles := []files.File{
		firstFile,
		secondFile,
	}

	//execute:
	rep := CreateFileRepositoryForTests()
	retFiles, retFilesErr := rep.RetrieveAll(retrieveInPath, fileNamesWithExt)

	if retFilesErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", retFilesErr.Error())
	}

	//verify:
	verifyFilesInRepositoryForTests(t, expectedFiles, retFiles)
}
