package chunks

import (
	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/chunks"
	concrete_files "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/files"
	concrete_hashtrees "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/hashtrees"
	concrete_stored_chunks "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/chunks"
)

// CreateChunksBuilderFactoryForTests creates a new ChunksBuilderFactory for tests
func CreateChunksBuilderFactoryForTests() chunk.ChunksBuilderFactory {
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	htBuilderFactory := concrete_hashtrees.CreateHashTreeBuilderFactoryForTests()
	chkSizeInBytes := 16
	extension := "chks"
	out := CreateChunksBuilderFactory(fileBuilderFactory, htBuilderFactory, chkSizeInBytes, extension)
	return out
}

// CreateChunksRepositoryForTests creates a new ChunksRepository for tests
func CreateChunksRepositoryForTests() chunk.ChunksRepository {
	htRepository := concrete_hashtrees.CreateHashTreeRepositoryForTests()
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	chksBuilderFactory := CreateChunksBuilderFactoryForTests()
	out := CreateChunksRepository(htRepository, fileRepository, chksBuilderFactory)
	return out
}

// CreateChunksServiceForTests creates a new ChunksService for tests
func CreateChunksServiceForTests() chunk.ChunksService {
	htService := concrete_hashtrees.CreateHashTreeServiceForTests()
	fileService := concrete_files.CreateFileServiceForTests()
	fileBuilderFactory := concrete_files.CreateFileBuilderFactoryForTests()
	storedChkBuilderFactory := concrete_stored_chunks.CreateBuilderFactoryForTests()
	out := CreateChunksService(htService, fileService, fileBuilderFactory, storedChkBuilderFactory)
	return out
}
