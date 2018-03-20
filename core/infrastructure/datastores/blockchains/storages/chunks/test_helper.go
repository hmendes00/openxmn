package chunks

import (
	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/datastores/blockchains/storages/chunks"
	concrete_file "github.com/XMNBlockchain/exmachina-network/core/infrastructure/datastores/blockchains/storages/files"
)

// CreateChunksForTests creates a Chunks for tests
func CreateChunksForTests() *Chunks {
	//variables:
	htFile := concrete_file.CreateFileForTests()
	chks := []*concrete_file.File{
		concrete_file.CreateFileForTests(),
		concrete_file.CreateFileForTests(),
		concrete_file.CreateFileForTests(),
	}

	out := createChunks(htFile, chks)
	return out.(*Chunks)
}

// CreateBuilderFactoryForTests creates a new ChunksBuilderFactory for tests
func CreateBuilderFactoryForTests() chunk.BuilderFactory {
	out := CreateBuilderFactory()
	return out
}

// CreateRepositoryForTests creates a Repository for tests
func CreateRepositoryForTests() chunk.Repository {
	fileRepository := concrete_file.CreateFileRepositoryForTests()
	chkBuilderFactory := CreateBuilderFactoryForTests()
	out := CreateRepository(fileRepository, chkBuilderFactory)
	return out
}
