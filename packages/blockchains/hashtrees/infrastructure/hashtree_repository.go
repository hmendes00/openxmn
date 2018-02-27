package infrastructure

import (
	"encoding/json"

	files "github.com/XMNBlockchain/core/packages/blockchains/files/domain"
	hashtrees "github.com/XMNBlockchain/core/packages/blockchains/hashtrees/domain"
)

// HashTreeRepository represents a concrete HashTree service
type HashTreeRepository struct {
	fileRepository files.FileRepository
}

// CreateHashTreeRepository creates a new HashTreeRepository instance
func CreateHashTreeRepository(fileRepository files.FileRepository) hashtrees.HashTreeRepository {
	out := HashTreeRepository{
		fileRepository: fileRepository,
	}

	return &out
}

// Retrieve retrieves an HashTree
func (rep *HashTreeRepository) Retrieve(dirPath string) (hashtrees.HashTree, error) {
	//read the hashtree:
	htFile, htFileErr := rep.fileRepository.Retrieve(dirPath, "hashtree.json")
	if htFileErr != nil {
		return nil, htFileErr
	}

	//unmarshal the hashtree:
	newHt := new(HashTree)
	jsonErr := json.Unmarshal(htFile.GetData(), newHt)
	if jsonErr != nil {
		return nil, jsonErr
	}

	//return the hashtree:
	return newHt, nil
}
