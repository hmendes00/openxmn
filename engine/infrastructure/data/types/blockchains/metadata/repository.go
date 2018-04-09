package metadata

import (
	"encoding/json"

	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
)

// Repository represents a concrete MetaData repository implementation
type Repository struct {
	fileRepository files.FileRepository
}

// CreateRepository creates a new Repository instance
func CreateRepository(fileRepository files.FileRepository) met.Repository {
	out := Repository{
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves a MetaData instance
func (rep *Repository) Retrieve(dirPath string) (met.MetaData, error) {
	//retrieve the file:
	fil, filErr := rep.fileRepository.Retrieve(dirPath, "metadata.json")
	if filErr != nil {
		return nil, filErr
	}

	//convert the js data into a MetaData instance:
	met := new(MetaData)
	jsErr := json.Unmarshal(fil.GetData(), met)
	if jsErr != nil {
		return nil, jsErr
	}

	return met, nil
}
