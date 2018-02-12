package infrastructure

import (
	"encoding/hex"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type fileService struct {
	storedFileBuilderFactory stored_files.FileBuilderFactory
}

// CreateFileService creates a new FileService instance
func CreateFileService(storedFileBuilderFactory stored_files.FileBuilderFactory) files.FileService {
	out := fileService{
		storedFileBuilderFactory: storedFileBuilderFactory,
	}
	return &out
}

// Save saves a file to disk
func (serv *fileService) Save(dirPath string, fil files.File) (stored_files.File, error) {

	//add the base path:
	toCreateDirPath := filepath.Join(dirPath, fil.GetDirPath())

	//create the directory recursively if the directory does not exists already:
	if _, err := os.Stat(toCreateDirPath); os.IsNotExist(err) {
		mkDirErr := os.MkdirAll(toCreateDirPath, os.ModePerm)
		if mkDirErr != nil {
			return nil, mkDirErr
		}
	}

	//create the file path:
	filePath := fil.GetFilePath()
	fullFilePath := filepath.Join(dirPath, filePath)

	//write the data on file:
	data := fil.GetData()
	wrErr := ioutil.WriteFile(fullFilePath, data, os.ModePerm)
	if wrErr != nil {
		return nil, wrErr
	}

	//build the stored file:
	ts := time.Now()
	h := fil.GetHash()
	hAsString := hex.EncodeToString(h.Sum(nil))
	sizeInBytes := fil.GetSizeInBytes()
	storedFile, storedFileErr := serv.storedFileBuilderFactory.
		Create().
		Create().
		WithPath(filePath).
		WithHash(hAsString).
		WithSizeInBytes(sizeInBytes).
		CreatedOn(ts).
		Now()

	if storedFileErr != nil {
		return nil, storedFileErr
	}

	return storedFile, nil
}

// SaveAll atomically saves multiple files to disk
func (serv *fileService) SaveAll(dirPath string, files []files.File) ([]stored_files.File, error) {
	out := []stored_files.File{}
	for _, oneFile := range files {
		oneStoredFile, oneStoredFileErr := serv.Save(dirPath, oneFile)
		if oneStoredFileErr != nil {
			return nil, oneStoredFileErr
		}
		out = append(out, oneStoredFile)
	}

	return out, nil
}

// Delete deletes a file from disk
func (serv *fileService) Delete(dirPath string, fileName string) error {
	fullPath := filepath.Join(dirPath, fileName)
	remErr := os.Remove(fullPath)
	if remErr != nil {
		return remErr
	}

	return nil
}

// Delete atomically deletes files from disk
func (serv *fileService) DeleteAll(dirPath string) error {
	remErr := os.RemoveAll(dirPath)
	if remErr != nil {
		return remErr
	}
	return nil
}
