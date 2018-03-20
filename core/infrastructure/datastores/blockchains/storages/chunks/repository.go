package chunks

import (
	"path/filepath"

	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/chunks"
	files "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/files"
)

// Repository represents a concrete stored chunks repository implementation
type Repository struct {
	fileRepository    files.FileRepository
	chkBuilderFactory chunk.BuilderFactory
}

// CreateRepository creates a Repository instance
func CreateRepository(fileRepository files.FileRepository, chkBuilderFactory chunk.BuilderFactory) chunk.Repository {
	out := Repository{
		fileRepository:    fileRepository,
		chkBuilderFactory: chkBuilderFactory,
	}

	return &out
}

// Retrieve retrieves a stored chunks instance
func (rep *Repository) Retrieve(dirPath string) (chunk.Chunks, error) {
	//hashtree:
	htFilePath := filepath.Join(dirPath, "hashtree.json")
	htFile, htFileErr := rep.fileRepository.Retrieve(htFilePath)
	if htFileErr != nil {
		return nil, htFileErr
	}

	//chunks:
	chksDirPath := filepath.Join(dirPath, "chunks")
	chksFiles, chksFilesErr := rep.fileRepository.RetrieveAll(chksDirPath)
	if chksFilesErr != nil {
		return nil, chksFilesErr
	}

	chks, chksErr := rep.chkBuilderFactory.Create().Create().WithHashTree(htFile).WithChunks(chksFiles).Now()
	if chksErr != nil {
		return nil, chksErr
	}

	return chks, nil
}
