package metadata

import (
	"strconv"
	"testing"
	"time"

	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/blockchains/metadata"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
	concrete_hashtrees "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/hashtrees"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
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

// CreateBuilderFactoryForTests creates a new BuilderFactory for tests
func CreateBuilderFactoryForTests() met.BuilderFactory {
	out := CreateBuilderFactory()
	return out
}

// CreateRepositoryForTests creates a new Repository for tests
func CreateRepositoryForTests() met.Repository {
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	out := CreateRepository(fileRepository)
	return out
}

// CreateServiceForTests creates a new Service for tests
func CreateServiceForTests() met.Service {
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactoryForTests()
	out := CreateService(fileBuilderFactory, fileService, storedFileBuilderFactory)
	return out
}
