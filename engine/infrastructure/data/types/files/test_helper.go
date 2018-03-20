package files

import (
	"crypto/sha256"

	files "github.com/XMNBlockchain/openxmn/engine/domain/data/types/files"
	concrete_stored_files "github.com/XMNBlockchain/openxmn/engine/infrastructure/data/stores/files"
)

// CreateFileForTests creates a File for tests
func CreateFileForTests() *File {
	//variables:
	extension := "tmp"
	fileName := "just_a_name"
	data := []byte("this is some data")
	sizeInBytes := len(data)
	h := sha256.New()
	h.Write(data)
	dirPath := ""

	out := createFile(h, sizeInBytes, data, dirPath, fileName, extension)
	return out.(*File)
}

// CreateFileBuilderFactoryForTests creates a new FileBuilderFactory for tests
func CreateFileBuilderFactoryForTests() files.FileBuilderFactory {
	out := CreateFileBuilderFactory()
	return out
}

// CreateFileRepositoryForTests creates a new FileRepository for tests
func CreateFileRepositoryForTests() files.FileRepository {
	fileBuilderFactory := CreateFileBuilderFactoryForTests()
	rep := CreateFileRepository(fileBuilderFactory)
	return rep
}

// CreateFileServiceForTests creates a new FileService for tests
func CreateFileServiceForTests() files.FileService {
	storedFileBuilderFactory := concrete_stored_files.CreateFileBuilderFactoryForTests()
	serv := CreateFileService(storedFileBuilderFactory)
	return serv
}
