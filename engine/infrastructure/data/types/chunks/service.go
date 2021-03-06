package chunks

import (
	"path/filepath"

	stored_chunks "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/chunks"
	chunk "github.com/XMNBlockchain/openxmn/engine/domain/data/types/chunks"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	hashtree "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// Service represents a concrete ChunksService implementation
type Service struct {
	htService               hashtree.HashTreeService
	fileService             files.FileService
	fileBuilderFactory      files.FileBuilderFactory
	storedChkBuilderFactory stored_chunks.BuilderFactory
}

// CreateService creates a new ChunksService instance
func CreateService(htService hashtree.HashTreeService, fileService files.FileService, fileBuilderFactory files.FileBuilderFactory, storedChkBuilderFactory stored_chunks.BuilderFactory) chunk.Service {
	out := Service{
		htService:               htService,
		fileService:             fileService,
		fileBuilderFactory:      fileBuilderFactory,
		storedChkBuilderFactory: storedChkBuilderFactory,
	}
	return &out
}

// Save saves a Chunks instance to disk
func (serv *Service) Save(dirPath string, chk chunk.Chunks) (stored_chunks.Chunks, error) {
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
	storedChk, storedChkErr := serv.storedChkBuilderFactory.Create().Create().WithChunks(storedFiles).WithHashTree(storedHt).Now()
	if storedChkErr != nil {
		return nil, storedChkErr
	}

	return storedChk, nil
}

// SaveAll saves []Chunks instances to disk
func (serv *Service) SaveAll(dirPath string, chks []chunk.Chunks) ([]stored_chunks.Chunks, error) {
	out := []stored_chunks.Chunks{}
	for _, oneChk := range chks {
		oneChkDirPath := filepath.Join(dirPath, oneChk.GetHashTree().GetHash().String())
		oneStoredChk, oneStoredChkErr := serv.Save(oneChkDirPath, oneChk)
		if oneStoredChkErr != nil {
			return nil, oneStoredChkErr
		}

		out = append(out, oneStoredChk)
	}

	return out, nil
}
