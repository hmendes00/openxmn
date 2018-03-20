package metadata

import (
	"strconv"
	"testing"
	"time"

	met "github.com/XMNBlockchain/exmachina-network/core/domain/data/types/blockchains/metadata"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/types/hashtrees"
	concrete_stored_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/data/stores/files"
	uuid "github.com/satori/go.uuid"
)

// JsDataForTests represents a structure for tests
type JsDataForTests struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateMetaDataForTests creates a MetaData for tests
func CreateMetaDataForTests(t *testing.T) *MetaData {
	//variables:
	id := uuid.NewV4()
	createdOn := time.Now().UTC()

	blocks := [][]byte{
		id.Bytes(),
		[]byte(strconv.Itoa(int(createdOn.UnixNano()))),
	}
	ht, _ := concrete_hashtrees.CreateHashTreeBuilderFactory().Create().Create().WithBlocks(blocks).Now()

	trs := createMetaData(&id, ht.(*concrete_hashtrees.HashTree), createdOn)
	return trs.(*MetaData)
}

// CreateMetaDataBuilderFactoryForTests creates a new MetaDataBuilderFactory for tests
func CreateMetaDataBuilderFactoryForTests() met.MetaDataBuilderFactory {
	out := CreateMetaDataBuilderFactory()
	return out
}

// CreateMetaDataRepositoryForTests creates a new MetaDataRepository for tests
func CreateMetaDataRepositoryForTests() met.MetaDataRepository {
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	out := CreateMetaDataRepository(fileRepository)
	return out
}

// CreateMetaDataServiceForTests creates a new MetaDataService for tests
func CreateMetaDataServiceForTests() met.MetaDataService {
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactoryForTests()
	out := CreateMetaDataService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	return out
}
