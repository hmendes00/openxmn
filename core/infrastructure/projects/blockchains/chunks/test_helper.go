package chunks

import (
	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/chunks"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/files"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/hashtrees"
	concrete_stored_chunks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/chunks"
)

// CreateBuilderFactoryForTests creates a new BuilderFactory for tests
func CreateBuilderFactoryForTests() chunk.BuilderFactory {
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	chkSizeInBytes := 16
	extension := "chks"
	out := CreateBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	return out
}

// CreateRepositoryForTests creates a new ChunksRepository for tests
func CreateRepositoryForTests() chunk.Repository {
	htRepository := concrete_hashtrees.CreateHashTreeRepositoryForTests()
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	chksBuilderFactory := CreateBuilderFactoryForTests()
	out := CreateRepository(htRepository, fileRepository, chksBuilderFactory)
	return out
}

// CreateServiceForTests creates a new ChunksService for tests
func CreateServiceForTests() chunk.Service {
	htService := concrete_hashtrees.CreateHashTreeServiceForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	storedChkBuilderFactory := concrete_stored_chunks.CreateBuilderFactoryForTests()
	out := CreateService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	return out
}
