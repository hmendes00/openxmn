package chained

import (
	"encoding/json"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	chained "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/blocks/validated/chained"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
)

// MetaDataService represents a concrete MetaData service implementation
type MetaDataService struct {
	fileBuilderFactory       files.FileBuilderFactory
	fileService              files.FileService
	storedFileBuilderFactory stored_files.FileBuilderFactory
}

// CreateMetaDataService creates a new MetaDataService instance
func CreateMetaDataService(fileBuilderFactory files.FileBuilderFactory, fileService files.FileService, storedFileBuilderFactory stored_files.FileBuilderFactory) chained.MetaDataService {
	out := MetaDataService{
		fileBuilderFactory:       fileBuilderFactory,
		fileService:              fileService,
		storedFileBuilderFactory: storedFileBuilderFactory,
	}

	return &out
}

// Save saves a MetaData instance
func (serv *MetaDataService) Save(dirPath string, met chained.MetaData) (stored_files.File, error) {
	//convert to json:
	js, jsErr := json.Marshal(met)
	if jsErr != nil {
		return nil, jsErr
	}

	//create the file:
	fil, filErr := serv.fileBuilderFactory.Create().Create().WithData(js).WithFileName("metadata").WithExtension("json").Now()
	if filErr != nil {
		return nil, filErr
	}

	//save the file:
	storedFile, storedFileErr := serv.fileService.Save(dirPath, fil)
	if storedFileErr != nil {
		return nil, storedFileErr
	}

	return storedFile, nil
}
