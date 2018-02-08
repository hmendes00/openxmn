package infrastructure

import (
	"testing"
	"time"

	concrete_file "github.com/XMNBlockchain/core/packages/storages/files/infrastructure"
)

// CreateChunksForTests creates a Chunks for tests
func CreateChunksForTests(t *testing.T) *Chunks {
	//variables:
	htFile := concrete_file.CreateFileForTests(t)
	chks := []*concrete_file.File{
		concrete_file.CreateFileForTests(t),
		concrete_file.CreateFileForTests(t),
		concrete_file.CreateFileForTests(t),
	}
	createdOn := time.Now().UTC()

	out := createChunks(htFile, chks, createdOn)
	return out.(*Chunks)
}
