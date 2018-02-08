package infrastructure

import (
	"crypto/sha256"
	"testing"
)

// CreateFileForTests creates a File for tests
func CreateFileForTests(t *testing.T) *File {
	//variables:
	extension := "tmp"
	data := []byte("this is some data")
	sizeInBytes := len(data)
	h := sha256.New()
	h.Write(data)

	out := createFile(h, sizeInBytes, data, extension)
	return out.(*File)
}
