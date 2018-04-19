package chunks

import (
	"io/ioutil"
	"os"
	"path/filepath"

	chunk "github.com/XMNBlockchain/openxmn/engine/domain/data/types/chunks"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	hashtree "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
)

// Repository represents a concrete ChunksRepository implementation
type Repository struct {
	htRepository       hashtree.HashTreeRepository
	fileRepository     files.FileRepository
	chksBuilderFactory chunk.BuilderFactory
}

// CreateRepository creates a new ChunksRepository instance
func CreateRepository(
	htRepository hashtree.HashTreeRepository,
	fileRepository files.FileRepository,
	chksBuilderFactory chunk.BuilderFactory,
) chunk.Repository {
	out := Repository{
		htRepository:       htRepository,
		fileRepository:     fileRepository,
		chksBuilderFactory: chksBuilderFactory,
	}
	return &out
}

// Retrieve retrieves a Chunks instance
func (rep *Repository) Retrieve(dirPath string) (chunk.Chunks, error) {

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
	ht, htErr := rep.htRepository.Retrieve(dirPath)
	if htErr != nil {
		return nil, htErr
	}

	//re-order the files data:
	filesData := [][]byte{}
	for _, oneFile := range files {
		filesData = append(filesData, oneFile.GetData())
	}

	orderedData, orderedErr := ht.Order(filesData)
	if orderedErr != nil {
		return nil, orderedErr
	}

	chks, chksErr := rep.chksBuilderFactory.Create().Create().WithBlocksData(orderedData).Now()
	if chksErr != nil {
		return nil, chksErr
	}

	return chks, nil
}

// RetrieveAll retrieves []Chunks instances
func (rep *Repository) RetrieveAll(dirPath string) ([]chunk.Chunks, error) {
	files, filesErr := ioutil.ReadDir(dirPath)
	if filesErr != nil {
		return nil, filesErr
	}

	chunks := []chunk.Chunks{}
	for _, oneFile := range files {
		if !oneFile.IsDir() {
			continue
		}

		chkDirPath := filepath.Join(dirPath, oneFile.Name())
		oneChk, oneChkErr := rep.Retrieve(chkDirPath)
		if oneChkErr != nil {
			return nil, oneChkErr
		}

		chunks = append(chunks, oneChk)
	}

	return chunks, nil
}
