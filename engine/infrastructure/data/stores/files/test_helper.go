package files

import (
	"time"

	dfil "github.com/XMNBlockchain/openxmn/engine/domain/data/stores/files"
)

// CreateFileForTests creates a File for tests
func CreateFileForTests() *File {
	//variables:
	path := "/tmp"
	data := []byte("this is some data")
	sizeInBytes := len(data)
	createdOn := time.Now().UTC()

	out := createFile(path, sizeInBytes, createdOn)
	return out.(*File)
}

// CreateFileBuilderFactoryForTests creates a new FileBuilderFactory for tests
func CreateFileBuilderFactoryForTests() dfil.FileBuilderFactory {
	out := CreateFileBuilderFactory()
	return out
}

// CreateFileRepositoryForTests creates a new FileRepository for tests
func CreateFileRepositoryForTests() dfil.FileRepository {
	fileBuilderFactory := CreateFileBuilderFactoryForTests()
	out := CreateFileRepository(fileBuilderFactory)
	return out
}
