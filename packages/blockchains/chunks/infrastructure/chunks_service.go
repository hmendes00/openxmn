package infrastructure

import (
	"path/filepath"
	"time"

	hashtree "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
	chunk "github.com/XMNBlockchain/core/packages/blockchains/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
	stored_chunks "github.com/XMNBlockchain/core/packages/storages/chunks/domain"
)

// ChunksService represents a concrete ChunksService implementation
type ChunksService struct {
	htService               hashtree.HashTreeService
	fileService             files.FileService
	fileBuilderFactory      files.FileBuilderFactory
	storedChkBuilderFactory stored_chunks.ChunksBuilderFactory
}

// CreateChunksService creates a new ChunksService instance
func CreateChunksService(htService hashtree.HashTreeService, fileService files.FileService, fileBuilderFactory files.FileBuilderFactory, storedChkBuilderFactory stored_chunks.ChunksBuilderFactory) chunk.ChunksService {
	out := ChunksService{
		htService:               htService,
		fileService:             fileService,
		fileBuilderFactory:      fileBuilderFactory,
		storedChkBuilderFactory: storedChkBuilderFactory,
	}
	return &out
}

// Save saves a Chunks instance to disk
func (serv *ChunksService) Save(dirPath string, chk chunk.Chunks) (stored_chunks.Chunks, error) {
	//save the hashtree:
	ht := chk.GetHashTree()
	storedHt, storedHtErr := serv.htService.Save(dirPath, ht)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	//save the chunks:
	files := chk.GetChunks()
	chksPath := filepath.Join(dirPath, "chunks")
	storedFiles, storedFilesErr := serv.fileService.SaveAll(chksPath, files)
	if storedFilesErr != nil {
		return nil, storedFilesErr
	}

	//build the stored chunks:
	ts := time.Now()
	storedChk, storedChkErr := serv.storedChkBuilderFactory.Create().Create().WithChunks(storedFiles).WithHashTree(storedHt).CreatedOn(ts).Now()
	if storedChkErr != nil {
		return nil, storedChkErr
	}

	return storedChk, nil
}
