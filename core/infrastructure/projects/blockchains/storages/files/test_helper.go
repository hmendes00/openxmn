package files

import (
	"crypto/sha256"
	"encoding/hex"
	"time"

	files "github.com/XMNBlockchain/exmachina-network/core/domain/projects/blockchains/storages/files"
)

// CreateFileForTests creates a File for tests
func CreateFileForTests() *File {
	//variables:
	path := "/tmp"
	data := []byte("this is some data")
	sizeInBytes := len(data)
	h := sha256.New()
	h.Write(data)
	hAsString := hex.EncodeToString(h.Sum(nil))
	createdOn := time.Now().UTC()

	out := createFile(path, hAsString, sizeInBytes, createdOn)
	return out.(*File)
}

// CreateFileBuilderFactoryForTests creates a new FileBuilderFactory for tests
func CreateFileBuilderFactoryForTests() files.FileBuilderFactory {
	out := CreateFileBuilderFactory()
	return out
}
