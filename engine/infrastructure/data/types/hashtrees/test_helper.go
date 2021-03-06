package hashtrees

import (
	"fmt"
	"math/rand"

	hashtrees "github.com/XMNBlockchain/openxmn/engine/domain/data/types/hashtrees"
	concrete_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/types/files"
)

// CreateHashTreeForTests creates an HashTree for tests
func CreateHashTreeForTests() *HashTree {
	//variables:
	r := rand.New(rand.NewSource(99))
	blks := [][]byte{
		[]byte("this"),
		[]byte("is"),
		[]byte("some"),
		[]byte("blocks"),
		[]byte(fmt.Sprintf("some rand number to make it unique: %d", r.Int())),
	}

	//execute:
	h, _ := createHashTreeFromBlocks(blks)
	return h.(*HashTree)
}

// CreateHashTreeBuilderFactoryForTests creates a new HashTreeBuilderFactory for tests
func CreateHashTreeBuilderFactoryForTests() hashtrees.HashTreeBuilderFactory {
	out := CreateHashTreeBuilderFactory()
	return out
}

// CreateHashTreeRepositoryForTests creates a new HashTreeRepository for tests
func CreateHashTreeRepositoryForTests() hashtrees.HashTreeRepository {
	fileRepository := concrete_files.CreateFileRepositoryForTests()
	out := CreateHashTreeRepository(fileRepository)
	return out
}

// CreateHashTreeServiceForTests creates a new HashTreeService for tests
func CreateHashTreeServiceForTests() hashtrees.HashTreeService {
	fileService := concrete_files.CreateFileServiceForTests()
	fileBuilderFactory := concrete_files.CreateFileBuilderFactory()
	out := CreateHashTreeService(fileService, fileBuilderFactory)
	return out
}
