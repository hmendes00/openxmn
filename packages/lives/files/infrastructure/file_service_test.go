package infrastructure

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
	concrete_stored_files "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
)

func createFileForTests(data []byte, fileName string, ext string) files.File {
	fil, _ := CreateFileBuilderFactory().Create().Create().WithData(data).WithDirPath("this_is_a_sub_dir").WithFileName(fileName).WithExtension(ext).Now()
	return fil
}

func verifyFilesInServiceForTests(t *testing.T, saveInPath string, fils []files.File, storedFiles []stored_files.File) {
	for index, oneFile := range fils {
		storedFile := storedFiles[index]
		retFirstHash := storedFile.GetHash()
		retFirstPath := storedFile.GetPath()
		retFirstSize := storedFile.GetSizeInBytes()

		h := sha256.New()
		h.Write(oneFile.GetData())
		hashAsString := hex.EncodeToString(h.Sum(nil))
		if hashAsString != retFirstHash {
			t.Errorf("the returned hash is invalid.  Expected: %s, Returned: %s", hashAsString, retFirstHash)
		}

		fullFilePath := oneFile.GetFilePath()
		if fullFilePath != retFirstPath {
			t.Errorf("the returned path is invalid.  Expected: %s, Returned: %s", fullFilePath, retFirstPath)
		}

		if oneFile.GetSizeInBytes() != retFirstSize {
			t.Errorf("the returned size is invalid.  Expected: %d, Returned: %d", oneFile.GetSizeInBytes(), retFirstSize)
		}

		//get the data from the file and compare it to the real data:
		toReadFilePath := filepath.Join(saveInPath, oneFile.GetFilePath())
		readContent, _ := ioutil.ReadFile(toReadFilePath)
		if !bytes.Equal(readContent, oneFile.GetData()) {
			t.Errorf("the written content is invalid.  Expected: %s, Read: %s", oneFile.GetData(), readContent)
		}
	}
}

func TestSave_withTextFile_thenDelete_Success(t *testing.T) {

	//file variables:
	saveInPath := filepath.Join("test_files", "files")
	files := []files.File{
		createFileForTests([]byte("this is some data"), "first_file", "txt"),
	}

	//service:
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	service := CreateFileService(storedFileBuilderFactory)

	//execute:
	storedFiles := []stored_files.File{}
	for _, oneFile := range files {
		storedFile, storedFileErr := service.Save(saveInPath, oneFile)
		if storedFileErr != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", storedFileErr.Error())
		}

		storedFiles = append(storedFiles, storedFile)
	}

	//verify:
	verifyFilesInServiceForTests(t, saveInPath, files, storedFiles)

	//delete:
	for _, oneFile := range files {
		//delete the file:
		toDelPath := filepath.Join(saveInPath, oneFile.GetDirPath())
		delErr := service.Delete(toDelPath, oneFile.GetFileNameWithExtension())
		if delErr != nil {
			t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
		}

		//verify that the file no longer exists:
		deletedFilePath := filepath.Join(toDelPath, oneFile.GetFileNameWithExtension())
		if _, err := os.Stat(deletedFilePath); err == nil {
			t.Errorf("the file (%s) should not exists", deletedFilePath)
		}

		//verify that the folder still exists:
		if _, err := os.Stat(saveInPath); os.IsNotExist(err) {
			t.Errorf("the directory (%s) should exists", saveInPath)
		}

		//delete the dir:
		delAllErr := service.DeleteAll(saveInPath)
		if delAllErr != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", delAllErr.Error())
		}

		//the directory should now no longer exists:
		if _, err := os.Stat(saveInPath); !os.IsNotExist(err) {
			t.Errorf("the directory was expected to be deleted: %s", saveInPath)
		}
	}
}

func TestSave_withBinaryFile_Success(t *testing.T) {

	//file variables:
	saveInPath := filepath.Join("test_files", "files")

	inputFilePath := filepath.Join("test_files", "input", "montreal.jpg")
	readData, readErr := ioutil.ReadFile(inputFilePath)
	if readErr != nil {
		t.Errorf("there was an error while reading the input binary file: %s", readErr.Error())
	}

	files := []files.File{
		createFileForTests(readData, "first_file", "jpg"),
	}

	//service:
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	service := CreateFileService(storedFileBuilderFactory)

	//execute:
	storedFiles := []stored_files.File{}
	for _, oneFile := range files {
		storedFile, storedFileErr := service.Save(saveInPath, oneFile)
		if storedFileErr != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", storedFileErr.Error())
		}

		storedFiles = append(storedFiles, storedFile)
	}

	//verify:
	verifyFilesInServiceForTests(t, saveInPath, files, storedFiles)

	//delete:
	for _, oneFile := range files {
		//delete the file:
		toDelPath := filepath.Join(saveInPath, oneFile.GetDirPath())
		delErr := service.Delete(toDelPath, oneFile.GetFileNameWithExtension())
		if delErr != nil {
			t.Errorf("the returned error was expected to be nil, error returned: %s", delErr.Error())
		}

		//verify that the file no longer exists:
		deletedFilePath := filepath.Join(toDelPath, oneFile.GetFileNameWithExtension())
		if _, err := os.Stat(deletedFilePath); err == nil {
			t.Errorf("the file (%s) should not exists", deletedFilePath)
		}

		//verify that the folder still exists:
		if _, err := os.Stat(saveInPath); os.IsNotExist(err) {
			t.Errorf("the directory (%s) should exists", saveInPath)
		}

		//delete the dir:
		delAllErr := service.DeleteAll(saveInPath)
		if delAllErr != nil {
			t.Errorf("the error was expected to be nil, error returned: %s", delAllErr.Error())
		}

		//the directory should now no longer exists:
		if _, err := os.Stat(saveInPath); !os.IsNotExist(err) {
			t.Errorf("the directory was expected to be deleted: %s", saveInPath)
		}
	}
}

func TestSaveAll_thenDeleteAll_Success(t *testing.T) {

	//file variables:
	saveInPath := filepath.Join("test_files", "files")

	//read the first file:
	firstInputFile := filepath.Join("test_files", "input", "montreal.jpg")
	firstReadData, firstReadDataErr := ioutil.ReadFile(firstInputFile)
	if firstReadDataErr != nil {
		t.Errorf("there was an error while reading the input binary file: %s", firstReadDataErr.Error())
	}

	//read the second file:
	secondInputFile := filepath.Join("test_files", "input", "montreal_second.jpg")
	secondReadData, secondReadDataErr := ioutil.ReadFile(secondInputFile)
	if secondReadDataErr != nil {
		t.Errorf("there was an error while reading the input binary file: %s", secondReadDataErr.Error())
	}

	files := []files.File{
		createFileForTests(firstReadData, "first_file", "jpg"),
		createFileForTests(secondReadData, "second_file", "jpg"),
	}

	//service:
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactory()
	service := CreateFileService(storedFileBuilderFactory)

	//execute:
	storedFiles, storedFilesErr := service.SaveAll(saveInPath, files)
	if storedFilesErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", storedFilesErr.Error())
	}

	//verify:
	verifyFilesInServiceForTests(t, saveInPath, files, storedFiles)

	//delete all:
	toDelErr := service.DeleteAll(saveInPath)
	if toDelErr != nil {
		t.Errorf("the returned error was expected to be nil, error returned: %s", toDelErr.Error())
	}

	//verify that the directory no longer exists:
	if _, toDelErr := os.Stat(saveInPath); toDelErr != nil {
		if !os.IsNotExist(toDelErr) {
			t.Errorf("the directory was expected to be deleted: %s.  The error was: %s", saveInPath, toDelErr.Error())
		}
	}
}
