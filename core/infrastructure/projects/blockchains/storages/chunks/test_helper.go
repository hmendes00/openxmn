package chunks

import (
	chunk "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/chunks"
	concrete_file "github.com/XMNBlockchain/exmachina-network/core/infrastructure/projects/blockchains/storages/files"
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
