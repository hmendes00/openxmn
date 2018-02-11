package infrastructure

import (
	"encoding/json"
	"path/filepath"
	"time"

	chunk "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
)

// ChunksService represents a concrete ChunksService implementation
type ChunksService struct {
	fileService             files.FileService
	fileBuilderFactory      files.FileBuilderFactory
	storedChkBuilderFactory stored_chunks.ChunksBuilderFactory
}

// CreateChunksService creates a new ChunksService instance
func CreateChunksService(fileService files.FileService, fileBuilderFactory files.FileBuilderFactory, storedChkBuilderFactory stored_chunks.ChunksBuilderFactory) chunk.ChunksService {
	out := ChunksService{
		fileService:             fileService,
		fileBuilderFactory:      fileBuilderFactory,
		storedChkBuilderFactory: storedChkBuilderFactory,
	}
	return &out
}

// Save saves a Chunks instance to disk
func (serv *ChunksService) Save(dirPath string, chk chunk.Chunks) (stored_chunks.Chunks, error) {

	//create the paths:
	h := chk.GetHashTree().GetHash().String()
	dirPathWithHash := filepath.Join(dirPath, h)

	//save the chunks:
	files := chk.GetChunks()
	storedFiles, storedFilesErr := serv.fileService.SaveAll(dirPathWithHash, files)
	if storedFilesErr != nil {
		return nil, storedFilesErr
	}

	//convert the hashtree to json:
	ht := chk.GetHashTree()
	js, jsErr := json.Marshal(ht)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the hashtree file:
	htFile, htFileErr := serv.fileBuilderFactory.Create().Create().WithData(js).WithExtension("json").Now()
	if htFileErr != nil {
		return nil, htFileErr
	}

	//save the hashtree:
	storedHt, storedHtErr := serv.fileService.Save(dirPathWithHash, htFile)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//build the stored chunks:
	ts := time.Now()
	storedChk, storedChkErr := serv.storedChkBuilderFactory.Create().Create().WithChunks(storedFiles).WithHashTree(storedHt).CreatedOn(ts).Now()
	if storedChkErr != nil {
		return nil, storedChkErr
	}

	return storedChk, nil
}

// Delete deletes a Chunks instance from disk
func (serv *ChunksService) Delete(dirPath string, hash string) error {
	dirPathWithHash := filepath.Join(dirPath, hash)
	delErr := serv.fileService.DeleteAll(dirPathWithHash)
	if delErr != nil {
		return delErr
	}

	return nil
}
