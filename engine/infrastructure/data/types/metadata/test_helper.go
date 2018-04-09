package metadata

import (
	"testing"
	"time"

	met "github.com/XMNBlockchain/openxmn/engine/domain/data/types/metadata"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
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
	lstUpOn := time.Now().UTC().Add(time.Second * 200 * -1)
	out := createMetaData(&id, createdOn, lstUpOn)
	return out.(*MetaData)
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
