package infrastructure

import (
	"encoding/json"

	files "github.com/XMNBlockchain/core/packages/lives/files/domain"
	objs "github.com/XMNBlockchain/core/packages/lives/objects/domain"
	stored_files "github.com/XMNBlockchain/core/packages/storages/files/domain"
)

// MetaDataService represents a concrete MetaData service implementation
type MetaDataService struct {
	fileBuilderFactory       files.FileBuilderFactory
	fileService              files.FileService
	storedFileBuilderFactory stored_files.FileBuilderFactory
}

// CreateMetaDataService creates a new MetaDataService instance
func CreateMetaDataService(fileBuilderFactory files.FileBuilderFactory, fileService files.FileService, storedFileBuilderFactory stored_files.FileBuilderFactory) objs.MetaDataService {
	out := MetaDataService{
		fileBuilderFactory:       fileBuilderFactory,
		fileService:              fileService,
		storedFileBuilderFactory: storedFileBuilderFactory,
	}

	return &out
}

// Save saves a MetaData instance
func (serv *MetaDataService) Save(dirPath string, met objs.MetaData) (stored_files.File, error) {
	//copnvert to json:
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
