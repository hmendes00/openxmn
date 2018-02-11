package infrastructure

import (
	"crypto/sha256"
	"testing"
)

// CreateFileForTests creates a File for tests
func CreateFileForTests(t *testing.T) *File {
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
