package hashtrees

import (
	"encoding/json"

	files "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/files"
	hashtrees "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/hashtrees"
	stored_file "github.com/XMNBlockchain/exmachina-network/core/domain/data/stores/files"
)

// HashTreeService represents a concrete HashTree service
type HashTreeService struct {
	fileService        files.FileService
	fileBuilderFactory files.FileBuilderFactory
}

// CreateHashTreeService creates a new HashTreeService instance
func CreateHashTreeService(fileService files.FileService, fileBuilderFactory files.FileBuilderFactory) hashtrees.HashTreeService {
	out := HashTreeService{
		fileService:        fileService,
		fileBuilderFactory: fileBuilderFactory,
	}

	return &out
}

// Save saves an HashTree
func (serv *HashTreeService) Save(dirPath string, ht hashtrees.HashTree) (stored_file.File, error) {
	//convert the hashtree to json:\
	js, jsErr := json.Marshal(ht)
	if jsErr != nil {
		return nil, jsErr
	}

	//build the hashtree file:
	htFile, htFileErr := serv.fileBuilderFactory.Create().Create().WithData(js).WithFileName("hashtree").WithExtension("json").Now()
	if htFileErr != nil {
		return nil, htFileErr
	}

	//save the hashtree:
	storedHt, storedHtErr := serv.fileService.Save(dirPath, htFile)
	if storedHtErr != nil {
		return nil, storedHtErr
	}

	return storedHt, nil
}
