package infrastructure

import (
	"encoding/gob"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"time"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

type fileService struct {
	storedFileBuilderFactory stored_files.FileBuilderFactory
	basePath                 string
}

// CreateFileService creates a new FileService instance
func CreateFileService(storedFileBuilderFactory stored_files.FileBuilderFactory, basePath string) files.FileService {
	out := fileService{
		storedFileBuilderFactory: storedFileBuilderFactory,
		basePath:                 basePath,
	}
	return &out
}

// Save saves a file to disk
func (serv *fileService) Save(dirPath string, fil files.File) (stored_files.File, error) {

	//add the base path:
	toCreateDirPath := filepath.Join(serv.basePath, dirPath)

	//create the directory recursively if the directory does not exists already:
	mkDirErr := os.MkdirAll(toCreateDirPath, os.ModePerm)
	if mkDirErr != nil {
		return nil, mkDirErr
	}

	//create the file path:
	ext := fil.GetExtension()
	hashAsString := hex.EncodeToString(fil.GetHash().Sum(nil))
	fileName := filepath.Join(hashAsString, ext)
	filePath := filepath.Join(dirPath, fileName)
	fullFilePath := filepath.Join(serv.basePath, filePath)

	//open the file:
	of, ofErr := os.Create(fullFilePath)
	if ofErr != nil {
		return nil, ofErr
	}
	defer of.Close()

	//write the data to disk:
	data := fil.GetData()
	encoder := gob.NewEncoder(of)
	encodeErr := encoder.Encode(data)
	if encodeErr != nil {
		return nil, encodeErr
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
	//create the tmp directory recursively:
	tmpDirPath := filepath.Join(dirPath, "/tmp")
	mkDirErr := os.MkdirAll(tmpDirPath, os.ModePerm)
	if mkDirErr != nil {
		return nil, mkDirErr
	}

	//save the files:
	tmpStoredFiles := []stored_files.File{}
	for _, oneFile := range files {
		oneStoredFile, oneStoredFileErr := serv.Save(tmpDirPath, oneFile)
		if oneStoredFileErr != nil {

			//delete the tmp directory
			remErr := serv.DeleteAll(tmpDirPath)
			if remErr != nil {
				return nil, remErr
			}

			return nil, oneStoredFileErr
		}

		tmpStoredFiles = append(tmpStoredFiles, oneStoredFile)
	}

	//build the new stored files according to the real dirPath:
	output := []stored_files.File{}
	for _, oneTmpStoredFile := range tmpStoredFiles {

		//get the data:
		h := oneTmpStoredFile.GetHash()
		sizeInBytes := oneTmpStoredFile.GetSizeInBytes()
		ts := oneTmpStoredFile.CreatedOn()

		//create the new file path:
		tmpFilePath := oneTmpStoredFile.GetPath()
		newFilePath := strings.Replace(tmpFilePath, tmpDirPath, dirPath, 1)
		oneStoredFile, oneStoredFileErr := serv.storedFileBuilderFactory.
			Create().
			Create().
			WithPath(newFilePath).
			WithHash(h).
			WithSizeInBytes(sizeInBytes).
			CreatedOn(ts).
			Now()

			//there was an error while building the new stored file, so delete the tmp directory and return the error:
		if oneStoredFileErr != nil {
			remErr := serv.DeleteAll(tmpDirPath)
			if remErr != nil {
				return nil, remErr
			}

			return nil, oneStoredFileErr
		}

		//add the new file to output:
		output = append(output, oneStoredFile)
	}

	//rename the tmp file to the real path:
	renErr := os.Rename(tmpDirPath, dirPath)
	if renErr != nil {
		return nil, renErr
	}

	return output, nil
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
