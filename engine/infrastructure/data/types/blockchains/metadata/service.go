package metadata

import (
	"encoding/json"

	stored_files "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
)

// Service represents a concrete MetaData service implementation
type Service struct {
	fileBuilderFactory       files.FileBuilderFactory
	fileService              files.FileService
	storedFileBuilderFactory stored_files.FileBuilderFactory
}

// CreateService creates a new Service instance
func CreateService(fileBuilderFactory files.FileBuilderFactory, fileService files.FileService, storedFileBuilderFactory stored_files.FileBuilderFactory) met.Service {
	out := Service{
		fileBuilderFactory:       fileBuilderFactory,
		fileService:              fileService,
		storedFileBuilderFactory: storedFileBuilderFactory,
	}

	return &out
}

// Save saves a MetaData instance
func (serv *Service) Save(dirPath string, met met.MetaData) (stored_files.File, error) {
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
