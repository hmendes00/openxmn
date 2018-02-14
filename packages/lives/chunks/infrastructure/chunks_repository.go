package infrastructure

import (
	"encoding/json"
	"os"
	"path/filepath"

	concrete_hashtrees "github.com/XMNBlockchain/core/packages/hashtrees/infrastructure"
	chunk "github.com/XMNBlockchain/core/packages/lives/chunks/domain"
	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
)

// ChunksRepository represents a concrete ChunksRepository implementation
type ChunksRepository struct {
	fileRepository     files.FileRepository
	chksBuilderFactory chunk.ChunksBuilderFactory
}

// CreateChunksRepository creates a new ChunksRepository instance
func CreateChunksRepository(fileRepository files.FileRepository, chksBuilderFactory chunk.ChunksBuilderFactory) chunk.ChunksRepository {
	out := ChunksRepository{
		fileRepository:     fileRepository,
		chksBuilderFactory: chksBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a Chunks instance
func (rep *ChunksRepository) Retrieve(dirPath string) (chunk.Chunks, error) {

	//create the paths:
	chksDirPath := filepath.Join(dirPath, "chunks")

	//make sure the dir exists:
	if _, err := os.Stat(chksDirPath); os.IsNotExist(err) {
		return nil, err
	}

	// scan the chunk file names:
	chkNames := []string{}
	walkErr := filepath.Walk(chksDirPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		chkNames = append(chkNames, info.Name())
		return nil
	})

	if walkErr != nil {
		return nil, walkErr
	}

	//retrieve the files:
	files, filesErr := rep.fileRepository.RetrieveAll(chksDirPath, chkNames)
	if filesErr != nil {
		return nil, filesErr
	}

	//read the hashtree:
	htFile, htFileErr := rep.fileRepository.Retrieve(dirPath, "hashtree.json")
	if htFileErr != nil {
		return nil, htFileErr
	}

	//unmarshal the hashtree:
	newHt := new(concrete_hashtrees.HashTree)
	jsonErr := json.Unmarshal(htFile.GetData(), newHt)
	if jsonErr != nil {
		return nil, jsonErr
	}

	//re-order the files data:
	filesData := [][]byte{}
	for _, oneFile := range files {
		filesData = append(filesData, oneFile.GetData())
	}

	orderedData, orderedErr := newHt.Order(filesData)
	if orderedErr != nil {
		return nil, orderedErr
	}

	chks, chksErr := rep.chksBuilderFactory.Create().Create().WithBlocksData(orderedData).Now()
	if chksErr != nil {
		return nil, chksErr
	}

	return chks, nil
}
