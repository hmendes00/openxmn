package infrastructure

import (
	"encoding/json"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
)

// MetaDataRepository represents a concrete MetaData repository implementation
type MetaDataRepository struct {
	fileRepository files.FileRepository
}

// CreateMetaDataRepository creates a new MetaDataRepository instance
func CreateMetaDataRepository(fileRepository files.FileRepository) objs.MetaDataRepository {
	out := MetaDataRepository{
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves a MetaData instance
func (rep *MetaDataRepository) Retrieve(dirPath string) (objs.MetaData, error) {
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
