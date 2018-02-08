package infrastructure

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"
)

// CreateFileForTests creates a File for tests
func CreateFileForTests(t *testing.T) *File {
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
